package money

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"exness/pkg/dto"
)

func NewTransferMoneyHandler(transferrer transferrer, transferValidator transferValidator) *transferMoneyHandler {
	return &transferMoneyHandler{
		transferrer: transferrer,
		validator:   transferValidator,
	}
}

type transferMoneyHandler struct {
	transferrer transferrer
	validator   transferValidator
}

func (h *transferMoneyHandler) Handle(ctx echo.Context) error {
	request := &dto.TransferMoneyRequest{}
	err := ctx.Bind(request)
	if err != nil {
		return err
	}

	err = h.validator.ValidateTransferRequest(request)
	if err != nil {
		return err
	}

	err = h.transferrer.TransferMoney(request.SenderAccountID, request.RecipientAccountID, request.Cents)
	if err != nil {
		return err
	}

	response := &dto.TransferMoneyResponse{
		Success: true,
	}
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, response)
}
