package config

import (
	"log"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT,required"`

	DatabaseURL string `env:"DATABASE_URL,required"`

	AppEnv string `env:"APP_ENV,required"`

	HOST string `env:"HOST,required"`

	AllowedOrigins string `env:"ALLOWED_ORIGINS" envDefault:"http://localhost:3000,http://127.0.0.1:3000,http://localhost:4173,http://127.0.0.1:4173,http://localhost:5173,http://127.0.0.1:5173"`
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

func (c *Config) GetAllowedOrigins() []string {
	rawOrigins := strings.Split(c.AllowedOrigins, ",")
	origins := make([]string, 0, len(rawOrigins))

	for _, origin := range rawOrigins {
		trimmedOrigin := strings.TrimSpace(origin)
		if trimmedOrigin == "" {
			continue
		}

		origins = append(origins, trimmedOrigin)
	}

	return origins
}
