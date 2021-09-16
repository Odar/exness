package money

import (
	"net/http"

	"exness/pkg/dto"

	"github.com/labstack/echo/v4"
)

func NewReplenishHandler(replenisher replenisher, replenishmentValidator replenishmentValidator) *replenishHandler {
	return &replenishHandler{
		replenisher: replenisher,
		validator:   replenishmentValidator,
	}
}

type replenishHandler struct {
	replenisher replenisher
	validator   replenishmentValidator
}

func (h *replenishHandler) Handle(ctx echo.Context) error {
	request := &dto.ReplenishRequest{}
	err := ctx.Bind(request)
	if err != nil {
		//TODO
		return ctx.JSON(http.StatusBadRequest, &dto.ApiError{
			Message: "invalid request",
		})
	}

	err = h.validator.ValidateReplenishRequest(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &dto.ApiError{
			Message: "invalid request",
		})
	}

	err = h.replenisher.ReplenishAccount(request.AccountID, request.Cents)
	if err != nil {
		//TODO
		return ctx.JSON(http.StatusInternalServerError, &dto.ApiError{
			Message: "...",
		})
	}

	response := &dto.ReplenishResponse{
		//TODO
		Cents: 0,
	}
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, response)
}
