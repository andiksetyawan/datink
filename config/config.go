package config

import (
	"log"

	"datink/pkg/config"
)

type Config struct {
	AppName     string `envconfig:"APP_NAME"`
	AppPort     int    `envconfig:"APP_PORT"`
	AppEnv      string `envconfig:"APP_ENV"`
	AppLogLevel string `envconfig:"APP_LOG_LEVEL"`

	DBName              string `envconfig:"DB_NAME"`
	DBUsername          string `envconfig:"DB_USERNAME"`
	DBPassword          string `envconfig:"DB_PASSWORD"`
	DBHost              string `envconfig:"DB_HOST"`
	DBPort              string `envconfig:"DB_PORT"`
	DBSSLMode           string `envconfig:"DB_SSL_MODE"`
	DBConnectionTimeout int    `envconfig:"DB_CONNECTION_TIMEOUT"`
	DBMaxIdleConns      int    `envconfig:"DB_MAX_IDLE_CONNS"`
	DBMaxOpenConns      int    `envconfig:"DB_MAX_OPEN_CONNS"`
	DBMaxLifetimeMinute int    `envconfig:"DB_MAX_LIFETIME_MINUTE"`

	JwtSigningKey string `envconfig:"JWT_SIGNING_KEY"`

	SwipeLimitFree int64 `envconfig:"SWIPE_LIMIT_FREE" default:"10"`
}

func NewConfig() (*Config, error) {
	configuration := Config{}
	if err := config.NewFromEnv(&configuration); err != nil {
		log.Fatal(err)
	}
	return &configuration, nil
}
