package account

import (
	"exness/pkg/dto"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func NewCreationHandler() *creationHandler {
	return &creationHandler{}
}

type creationHandler struct{}

func (h *creationHandler) Handle(ctx echo.Context) error {
	request := &dto.AccountCreationRequest{}
	err := ctx.Bind(request)
	if err != nil {
		return errors.Wrap(err, "can not bind to api struct")
	}

}
