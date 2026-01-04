package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	DatabaseURL    string
	Port           string
	Environment    string
	OpenAIAPIKey   string
	AWSRegion      string
	AWSBucket      string
	JWTSecret      string
	RedisURL       string
	AllowedOrigins []string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://localhost:5432/wat2do?sslmode=disable"),
		Port:           getEnv("PORT", "8000"),
		Environment:    getEnv("ENVIRONMENT", "development"),
		OpenAIAPIKey:   getEnv("OPENAI_API_KEY", ""),
		AWSRegion:      getEnv("AWS_REGION", "us-east-1"),
		AWSBucket:      getEnv("AWS_BUCKET", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		RedisURL:       getEnv("REDIS_URL", "localhost:6379"),
		AllowedOrigins: []string{getEnv("ALLOWED_ORIGINS", "*")},
	}

	return config
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
