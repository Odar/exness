package main

import (
	"exness/internal/config"
	"fmt"
	"os"
)

func main() {
	cfg := initConfig()

	fmt.Printf(`-dir=%s postgres "%s" up`, cfg.ExnessDB.MigrationDir, cfg.ExnessDB.DNS())
}

func initConfig() config.Options {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cfg
}
