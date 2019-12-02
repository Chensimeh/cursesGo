package apiserver

import "github.com/trzhensimekh/cursesGo/task2Rest/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	Store *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		Store: store.NewConfig(),
	}
}
