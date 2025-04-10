package config

import (
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required" envDefault:"3000"`
	DBHost     string `env:"DB_HOST,required" envDefault:"localhost"`
	DBName     string `env:"DB_NAME,required" envDefault:"event_manager"`
	DBUser     string `env:"DB_USER,required" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD,required" envDefault:"postgres"`
	DBSSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file at config/config.go: %v", err)
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Error parsing env config at config/config.go: %v", err)
	}

	return config

}
