package db

import "fmt"

type Config struct {
	Addr         string
	Port         uint16
	User         string
	Password     string
	DB           string
	MigrationDir string
}

func (c *Config) DNS() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", c.User, c.Password, c.Addr, c.Port, c.DB)
}
