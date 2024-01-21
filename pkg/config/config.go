package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func NewFromEnv(object interface{}) error {
	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", object); err != nil {
			return fmt.Errorf("failed to read from env variable: %w", err)
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return fmt.Errorf("failed to read from .env file: %w", err)
	}

	if err := envconfig.Process("", object); err != nil {
		return fmt.Errorf("failed to read from env variable: %w", err)
	}

	return nil
}
