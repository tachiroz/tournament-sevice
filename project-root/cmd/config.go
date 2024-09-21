package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	JWTSecret  string
}

func LoadConfig() Config {
	return Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "tournament_db"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecretkey"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBName)
}
