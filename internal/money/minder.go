package money

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"exness/internal/models"
)

func NewMinder(postgres *sqlx.DB) *minder {
	return &minder{
		postgres: postgres,
		builder:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

type replenisher interface {
	ReplenishAccount(accountID int64, cents int64) error
}

type transferrer interface {
	TransferMoney(fromAccountID int64, toAccountID int64, cents int64) error
}

type minder struct {
	postgres *sqlx.DB
	builder  sq.StatementBuilderType
}

func (m *minder) ReplenishAccount(accountID int64, cents int64) (err error) {
	tx, err := m.postgres.BeginTxx(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "can not start transaction")
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = errors.Wrapf(err, "error during rollback (err: %s)", rollbackErr)
				return
			}

			return
		}
	}()

	err = m.replenishAccount(tx, accountID, cents)
	if err != nil {
		return errors.Wrap(err, "can not replenish account")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "can not commit transaction")
	}

	return nil
}

func (m *minder) TransferMoney(fromAccountID int64, toAccountID int64, cents int64) error {
	tx, err := m.postgres.BeginTxx(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "can not start transaction")
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = errors.Wrapf(err, "error during rollback (err: %s)", rollbackErr)
				return
			}

			return
		}
	}()

	err = m.transferMoney(tx, fromAccountID, toAccountID, cents)
	if err != nil {
		return errors.Wrap(err, "can not transfer money")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "can not commit transaction")
	}

	return nil
}

func (m *minder) replenishAccount(tx *sqlx.Tx, accountID int64, cents int64) error {
	accounts, err := m.getAndLockAccounts(tx, accountID)
	if err != nil {
		return errors.Wrap(err, "can not get account")
	}
	if len(accounts) == 0 {
		return errors.New("can not find account by id")
	}
	if len(accounts) > 1 {
		return errors.New("found more account than one")
	}

	account := accounts[0]
	if isSumOverflow(cents, account.Balance) {
		return errors.New("cents will be overflowing")
	}

	transaction := models.Transaction{
		RecipientAccountID: accountID,
		Cents:              cents,
		Type:               models.ReplenishTransactionType,
	}
	err = m.addTransaction(tx, transaction)
	if err != nil {
		return errors.Wrap(err, "can not add transaction")
	}

	err = m.addMoneyToAccount(tx, accountID, account.Balance+cents)
	if err != nil {
		return errors.Wrap(err, "can not add cents to account")
	}

	return nil
}

func (m *minder) transferMoney(tx *sqlx.Tx, senderAccountID int64, recipientAccountID int64, cents int64) error {
	accounts, err := m.getAndLockAccounts(tx, senderAccountID, recipientAccountID)
	if err != nil {
		return errors.Wrap(err, "can not get account")
	}
	if len(accounts) > 2 {
		return errors.New("found more accounts then two")
	}

	var (
		senderAccount    *models.Account
		recipientAccount *models.Account
	)
	for _, account := range accounts {
		if account.ID == senderAccountID {
			senderAccount = account
			continue
		}
		if account.ID == recipientAccountID {
			recipientAccount = account
			continue
		}
	}
	if senderAccount == nil {
		return errors.New("can't find an account for cents withdrawal")
	}
	if recipientAccount == nil {
		return errors.New("can't find an account to accrue cents")
	}

	if isSumOverflow(cents, recipientAccount.Balance) {
		return errors.New("cents will be overflowing")
	}
	if isBelowZero(senderAccount.Balance, cents) {
		return errors.New("balance cannot be negative")
	}

	transaction := models.Transaction{
		SenderAccountID:    senderAccountID,
		RecipientAccountID: recipientAccountID,
		Cents:              cents,
		Type:               models.TransferTransactionType,
	}
	err = m.addTransaction(tx, transaction)
	if err != nil {
		return errors.Wrap(err, "can not add transaction")
	}

	err = m.addMoneyToAccount(tx, senderAccountID, senderAccount.Balance-cents)
	if err != nil {
		return errors.Wrap(err, "can not add cents to account")
	}

	err = m.addMoneyToAccount(tx, recipientAccountID, recipientAccount.Balance+cents)
	if err != nil {
		return errors.Wrap(err, "can not add cents to account")
	}

	return nil
}

func (m *minder) getAndLockAccounts(tx *sqlx.Tx, accountIDs ...int64) ([]*models.Account, error) {
	builder := m.builder.
		Select("id, balance").
		From("account").
		Where(sq.Eq{"id": accountIDs}).
		Suffix("FOR UPDATE")
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "can not build sql")
	}

	accounts := make([]*models.Account, 0)
	err = tx.Select(&accounts, query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "can not exec sql '%s' with args '%+v'", query, args)
	}

	return accounts, nil
}

func (m *minder) addTransaction(tx *sqlx.Tx, transaction models.Transaction) error {
	builder := m.builder.
		Insert("transaction").
		Columns("sender_account_id", "recipient_account_id", "cents", "type").
		Values(transaction.SenderAccountID, transaction.RecipientAccountID, transaction.Cents, transaction.Type)
	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "can not build sql")
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return errors.Wrapf(err, "can not exec sql '%s' with args '%+v'", query, args)
	}

	return nil
}

func (m *minder) addMoneyToAccount(tx *sqlx.Tx, accountID int64, newBalance int64) error {
	builder := m.builder.
		Update("account").
		Set("balance", newBalance).
		Where("id = ?", accountID)
	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "can not build sql")
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return errors.Wrapf(err, "can not exec sql '%s' with args '%+v'", query, args)
	}

	return nil
}
