package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HttpPort               string   `env:"HTTP_PORT"                 env-required:"false" env-default:"8000"`
	PostgresDSN            string   `env:"POSTGRES_DSN"              env-required:"true"  env-default:"postgres://postgres:password@localhost:5432/main?sslmode=disable"`
	MemcachedAddrs         []string `env:"MEMCACHED_ADDRS"           env-required:"true"  env-default:"localhost:11211"`
	S3Endpoint             string   `env:"S3_ENDPOINT"               env-required:"true"  env-default:"TODO"`
	S3AccessKey            string   `env:"S3_ACCESS_KEY"             env-required:"true"  env-default:"TODO"`
	S3SecretKey            string   `env:"S3_SECRET_KEY"             env-required:"true"  env-default:"TODO"`
	S3UseSSL               bool     `env:"S3_USE_SSL"                env-required:"true"  env-default:"false"`
	P1SMSApiKey            string   `env:"P1SMS_API_KEY"             env-required:"true"  env-default:"TODO"`
	SMTPHostname           string   `env:"SMTP_HOSTNAME"             env-required:"true"  env-default:"TODO"`
	SMTPPort               int      `env:"SMTP_PORT"                 env-required:"true"  env-default:"587"`
	SMTPUsername           string   `env:"SMTP_USERNAME"             env-required:"true"  env-default:"TODO"`
	SMTPPassword           string   `env:"SMTP_PASSWORD"             env-required:"true"  env-default:"TODO"`
	DomainName             string   `env:"DOMAIN_NAME"               env-required:"true"  env-default:"TODO"`
	DomainSecure           bool     `env:"DOMAIN_SECURE"             env-required:"true"  env-default:"false"`
	CustomerAuthCookieName string   `env:"CUSTOMER_AUTH_COOKIE_NAME" env-required:"true"  env-default:"TODO"`
	UserAuthCookieName     string   `env:"USER_AUTH_COOKIE_NAME"     env-required:"true"  env-default:"TODO"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return cfg, nil
}
