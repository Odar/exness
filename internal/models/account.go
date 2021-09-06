package models

import "time"

type Account struct {
	ID        int64
	Cents     int64
	CreatedAt time.Time
}
