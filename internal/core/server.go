package core

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(cfg Config, handlers []*Handler) *server {
	e := echo.New()
	e.HideBanner = true

	for _, handler := range handlers {
		e.Add(handler.Method, handler.Path, handler.HandlerFunc.Handle, middleware.Recover())
	}

	return &server{
		echo:     e,
		handlers: handlers,
		address:  fmt.Sprintf(":%d", cfg.Port),
	}
}

type server struct {
	echo     *echo.Echo
	handlers []*Handler
	address  string
}

func (s *server) Setup() {
	for _, handler := range s.handlers {
		s.echo.Add(handler.Method, handler.Path, handler.HandlerFunc.Handle, middleware.Recover())
	}
}

func (s *server) Start() error {
	return s.echo.Start(s.address)
}

func (s *server) Stop() error {
	return s.echo.Shutdown(context.Background())
}
