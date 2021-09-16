package config

import (
	"exness/internal/core"
	"exness/internal/db"
)

type Options struct {
	ExnessDB db.Config
	Server   core.Config
	Logger   Logger
}

type Logger struct {
	Level string
}

var defaultOptions = Options{}
