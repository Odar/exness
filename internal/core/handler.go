package core

import "github.com/labstack/echo/v4"

//TODO naming
type HandlerStruct interface {
	Handle(ctx echo.Context) error
}

type Handler struct {
	Method      string
	Path        string
	HandlerFunc HandlerStruct
}
