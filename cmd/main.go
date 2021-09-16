package main

import (
	"fmt"

	"exness/internal/account"
	"exness/internal/config"
	"exness/internal/core"
	"exness/internal/db"
	"exness/internal/money"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)

	postgres, err := db.NewPostgres(cfg.ExnessDB)
	if err != nil {
		panic(err)
	}

	moneyValidator := money.NewValidator()
	moneyMinder := money.NewMinder(postgres)
	accountCreator := account.NewCreator(postgres)

	replenishHandler := money.NewReplenishHandler(moneyMinder, moneyValidator)
	transferHandler := money.NewTransferMoneyHandler(moneyMinder, moneyValidator)
	creatorHandler := account.NewCreationHandler(accountCreator)

	handlers := prepareHandlers(replenishHandler, transferHandler, creatorHandler)
	server := core.NewServer(cfg.Server, handlers)

	server.Start()
}

func prepareHandlers(
	replenishHandler core.HandlerStruct,
	transferHandler core.HandlerStruct,
	creatorHandler core.HandlerStruct,
) []*core.Handler {
	return []*core.Handler{
		{
			Method:      "PUT",
			Path:        "/account",
			HandlerFunc: creatorHandler,
		},
		{
			Method:      "POST",
			Path:        "/account/replenish",
			HandlerFunc: replenishHandler,
		},
		{
			Method:      "POST",
			Path:        "/account/transfer",
			HandlerFunc: transferHandler,
		},
	}
}
