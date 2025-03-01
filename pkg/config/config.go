package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all configuration variables
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	ServerPort string
}

// LoadConfig loads environment variables from .env file
func LoadConfig() *Config {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBName:     getEnv("DB_NAME", "tasks_db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return config
}

// getEnv retrieves environment variables with a fallback default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
