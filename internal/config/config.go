package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL       string
	Port              string
	OxeliaGatewayMode bool
}

var Cfg *Config

func Load() *Config {
	_ = godotenv.Load()

	Cfg = &Config{
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		Port:              getEnv("AIHELPER_PORT", "8004"),
		OxeliaGatewayMode: getEnvBool("OXELIA_GATEWAY_MODE", false),
	}

	if Cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return Cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return b
}
