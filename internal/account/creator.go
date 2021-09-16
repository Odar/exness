package account

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"exness/internal/models"
)

func NewCreator(postgres *sqlx.DB) *creator {
	return &creator{
		postgres: postgres,
		builder:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

type creatorI interface {
	CreateAccount() (*models.Account, error)
}

type creator struct {
	postgres *sqlx.DB
	builder  sq.StatementBuilderType
}

func (c *creator) CreateAccount() (*models.Account, error) {
	builder := c.builder.
		Insert("account").
		Columns("balance").
		Values(0).
		Suffix("RETURNING id")
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "can not build sql")
	}

	var id int64
	err = c.postgres.Get(&id, query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "can not exec sql '%s' with args '%+v'", query, args)
	}

	return &models.Account{
		ID: id,
	}, nil
}
