package money

import (
	"exness/pkg/dto"

	"github.com/labstack/echo/v4"
)

func NewReplenishHandler(replenisher replenisher) *replenishHandler {
	return &replenishHandler{
		replenisher: replenisher,
	}
}

type replenishHandler struct {
	replenisher replenisher
}

func (h *replenishHandler) Handle(ctx echo.Context) error {
	request := &dto.ReplenishRequest{}
	err := ctx.Bind(request)
	if err != nil {

	}

	err = h.replenisher.ReplenishAccount(request.AccountID, request.Cents)
	if err != nil {

	}

	return nil
}
