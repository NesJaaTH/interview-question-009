package config

import (
	"log"
	"os"

	"example.com/interview-question-009/pkg"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	AppEnv      string
	CORSOrigins []string
	JWTSecret   string
}

func Load(envFile string) *Config {
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("warn: could not load %s, falling back to environment variables", envFile)
	}

	return &Config{
		AppPort:     getEnv("APP_PORT", "7809"),
		AppEnv:      getEnv("APP_ENV", "development"),
		CORSOrigins: pkg.ParseStringList(getEnv("CORS_ORIGINS", "*")),
		JWTSecret:   getEnv("JWT_SECRET", "dev-secret-change-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// IsProduction reports whether the app is running in production mode.
func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}

// IsDevelopment reports whether the app is running in development mode.
func (c *Config) IsDevelopment() bool {
	return c.AppEnv == "development"
}
