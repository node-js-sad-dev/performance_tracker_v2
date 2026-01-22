package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT,required"`

	DatabaseURL string `env:"DATABASE_URL,required"`

	AppEnv string `env:"APP_ENV,required"`

	HOST string `env:"HOST,required"`
}

func Load() *Config {
	_ = godotenv.Load("./config/.env")

	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Config validation failed: %v", err)
	}

	return &cfg
}

func (c *Config) GetServerAddr() string {
	return ":" + c.Port
}
