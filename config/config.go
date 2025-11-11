package config

import (
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
)

type Config struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func Load() *Config {
	config := &Config{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}
	return config
}
