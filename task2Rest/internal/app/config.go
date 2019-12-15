package app

import (
	"github.com/trzhensimekh/cursesGo/task2Rest/internal/data"
)

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	Store       *data.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		Store:    data.NewConfig(),
	}
}
