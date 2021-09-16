package models

import "time"

type Account struct {
	ID        int64     `db:"id"`
	Balance   int64     `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
}
