package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg Config) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", cfg.DNS())
}
