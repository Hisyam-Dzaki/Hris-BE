package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppPort       string
	AppEnv        string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	JWTSecret     string
	JWTExpiryHour int
}

func Load() *Config {
	expiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))

	return &Config{
		AppPort:       getEnv("APP_PORT", "8080"),
		AppEnv:        getEnv("APP_ENV", "development"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "hris_admin"),
		DBPassword:    getEnv("DB_PASSWORD", "hris_secret_2026"),
		DBName:        getEnv("DB_NAME", "hris_db"),
		DBSSLMode:     getEnv("DB_SSLMODE", "disable"),
		JWTSecret:     getEnv("JWT_SECRET", "hris-jwt-secret-change-in-production"),
		JWTExpiryHour: expiry,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
