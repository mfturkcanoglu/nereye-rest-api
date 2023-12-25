package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	logger *log.Logger
}

func NewConfig(logger *log.Logger) *Config {
	return &Config{
		logger: logger,
	}
}

func (config *Config) LoadEnv() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		config.logger.Panicln(
			"Error occured while loading env variables",
			err,
		)
	}

	return config
}
