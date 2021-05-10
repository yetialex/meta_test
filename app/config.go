package app

import (
	"errors"
	"os"
)

type Config struct {
	LoggingLevel string
	DatabaseURL  string
}

func NewConfig() (*Config, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, errors.New("db url is empty")
	}

	return &Config{
		DatabaseURL:  dbURL,
		LoggingLevel: os.Getenv("LOGGING_LEVEL"),
	}, nil
}
