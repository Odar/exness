package money

import (
	"github.com/pkg/errors"

	"exness/internal/core"
	"exness/pkg/dto"
)

func NewValidator() *validator {
	return &validator{}
}

type replenishmentValidator interface {
	ValidateReplenishRequest(request *dto.ReplenishRequest) error
}

type transferValidator interface {
	ValidateTransferRequest(request *dto.TransferMoneyRequest) error
}

type validator struct {
}

func (v *validator) ValidateReplenishRequest(request *dto.ReplenishRequest) error {
	errs := make([]error, 0)
	if request.AccountID <= 0 {
		errs = append(errs, errors.New("account id cannot be zero or below"))
	}

	if request.Cents <= 0 {
		errs = append(errs, errors.New("cents can not be zero or below"))
	}

	return core.CollapseErrors(errs)
}

func (v *validator) ValidateTransferRequest(request *dto.TransferMoneyRequest) error {
	errs := make([]error, 0)
	if request.SenderAccountID <= 0 {
		errs = append(errs, errors.New("sender account id cannot be zero or below"))
	}

	if request.RecipientAccountID <= 0 {
		errs = append(errs, errors.New("recipient account id cannot be zero or below"))
	}

	if request.Cents <= 0 {
		errs = append(errs, errors.New("cents can not be zero or below"))
	}

	return core.CollapseErrors(errs)
}
