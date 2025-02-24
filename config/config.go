package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PostgresDSN string  `env:"POSTGRES_DSN" env-required:"true" env-default:"postgres://postgres:password@localhost:5432/main?sslmode=disable"`
	BotToken    string  `env:"BOT_TOKEN"    env-required:"true"`
	AdminIDs    []int64 `env:"ADMIN_IDS"    env-required:"true"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return cfg, nil
}
