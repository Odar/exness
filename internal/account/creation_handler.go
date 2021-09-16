package account

import (
	"net/http"

	"exness/pkg/dto"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func NewCreationHandler(creator creatorI) *creationHandler {
	return &creationHandler{
		creator: creator,
	}
}

type creationHandler struct {
	creator creatorI
}

func (h *creationHandler) Handle(ctx echo.Context) error {
	newAccount, err := h.creator.CreateAccount()
	if err != nil {
		return errors.Wrap(err, "can not create account")
	}

	response := &dto.AccountCreationResponse{
		Account: &dto.Account{
			ID:    newAccount.ID,
			Cents: newAccount.Balance,
		},
	}
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, response)
}
