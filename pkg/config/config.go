package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
}

type AppConfig struct {
	Host string
	Port string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	sslMode, err := strconv.ParseBool(os.Getenv("POSTGRES_SSL_MODE"))
	if err != nil {
		return nil, err
	}
	return &Config{
		App: AppConfig{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		Postgres: PostgresConfig{
			PostgresqlHost:     os.Getenv("POSTGRES_DB_HOST"),
			PostgresqlPort:     os.Getenv("POSTGRES_DB_PORT"),
			PostgresqlUser:     os.Getenv("POSTGRES_DB_USER"),
			PostgresqlPassword: os.Getenv("POSTGRES_DB_PASSWORD"),
			PostgresqlDbname:   os.Getenv("POSTGRES_DB"),
			PostgresqlSSLMode:  sslMode,
		},
	}, nil
}
