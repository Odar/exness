package models

import "time"

type Transaction struct {
	ID                 int64           `db:"id"`
	SenderAccountID    int64           `db:"sender_account_id"`
	RecipientAccountID int64           `db:"recipient_account_id"`
	Cents              int64           `db:"cents"`
	Type               TransactionType `db:"transaction_type"`
	CommittedIn        time.Time       `db:"committed_in"`
}

type TransactionType string

const (
	ReplenishTransactionType TransactionType = "replenish"
	TransferTransactionType                  = "transfer"
)
