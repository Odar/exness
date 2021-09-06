package models

import "time"

type Transaction struct {
	ID          int64
	From        int64
	To          int64
	Cents       int64
	Type        TransactionType
	CommittedIn time.Time
}

type TransactionType string

const (
	ReplenishTransactionType TransactionType = "replenish"
	TransferTransactionType                  = "transfer"
)
