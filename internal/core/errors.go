package core

import (
	"strings"

	"github.com/pkg/errors"

	"exness/pkg/dto"
)

func CollapseErrors(errs []error) error {
	messages := make([]string, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			messages = append(messages, err.Error())
		}
	}

	if len(messages) == 0 {
		return nil
	}

	return errors.Errorf("collapsed errs: %s", strings.Join(messages, "; "))
}

func NewClientError(err error) error {
	return &clientError{
		err: err,
	}
}

type clientError struct {
	err           error
	clientMessage string
}

func (e *clientError) Error() string {
	return e.err.Error()
}

func (e *clientError) SetClientMsg(msg string) *clientError {
	if e == nil {
		return nil
	}

	e.clientMessage = msg

	return e
}

func (e *clientError) AddToClientMsg(x string) *clientError {
	if e == nil {
		return nil
	}

	e.clientMessage = e.clientMessage + x

	return e
}

func (e *clientError) ToApiError() *dto.ApiError {
	if e == nil {
		return nil
	}

	if e.clientMessage == "" {
		return &dto.ApiError{
			Message: "Internal server error",
		}
	}

	return &dto.ApiError{
		Message: e.clientMessage,
	}
}
