package config

import "exness/internal/db"

type Options struct {
	ExnessDB db.Config
	Logger   Logger
}

type Logger struct {
	Level string
}

var defaultOptions = Options{}
