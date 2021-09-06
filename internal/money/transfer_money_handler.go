package money

import (
	"github.com/labstack/echo/v4"

	"exness/pkg/dto"
)

func NewTransferMoneyHandler(transferrer transferrer) *transferMoneyHandler {
	return &transferMoneyHandler{
		transferrer: transferrer,
	}
}

type transferMoneyHandler struct {
	transferrer transferrer
}

func (h *transferMoneyHandler) Handle(ctx echo.Context) error {
	request := &dto.TransferMoneyRequest{}
	err := ctx.Bind(request)
	if err != nil {

	}

	err = h.transferrer.TransferMoney(request.FromAccountID, request.ToAccountID, request.Cents)

	return nil
}
