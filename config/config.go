package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig
	App      AppConfig
}

type AppConfig struct {
	AppVersion        string
	Mode              string
	Port              string
	SSL               bool
	CSRF              bool
	Debug             bool
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlPort     string
	PostgresqlSSLMode  bool
}

func InitConfig(path, filename string) (*Config, error) {
	// Load Config from file
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(path)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}

	cfg := Config{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &cfg, err
}
