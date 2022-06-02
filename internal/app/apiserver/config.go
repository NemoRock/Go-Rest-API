package apiserver

import "Go_REST_API/internal/app/store"

type Config struct {
	BimAddr  string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BimAddr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
