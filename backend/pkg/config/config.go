package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	DB       DBConfig
	JWT      JWTConfig
	CORS     CORSConfig
}

type AppConfig struct {
	Env  string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	AccessSecret        string
	RefreshSecret       string
	AccessExpiryMinutes int
	RefreshExpiryDays   int
}

type CORSConfig struct {
	AllowedOrigin string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	accessExpiry, err := strconv.Atoi(getEnv("JWT_ACCESS_EXPIRY_MINUTES", "15"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_ACCESS_EXPIRY_MINUTES: %w", err)
	}

	refreshExpiry, err := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRY_DAYS", "7"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_REFRESH_EXPIRY_DAYS: %w", err)
	}

	cfg := &Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		DB: DBConfig{
			Host:     mustEnv("POSTGRES_HOST"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     mustEnv("POSTGRES_USER"),
			Password: mustEnv("POSTGRES_PASSWORD"),
			Name:     mustEnv("POSTGRES_DB"),
		},
		JWT: JWTConfig{
			AccessSecret:        mustEnv("JWT_ACCESS_SECRET"),
			RefreshSecret:       mustEnv("JWT_REFRESH_SECRET"),
			AccessExpiryMinutes: accessExpiry,
			RefreshExpiryDays:   refreshExpiry,
		},
		CORS: CORSConfig{
			AllowedOrigin: getEnv("CORS_ALLOWED_ORIGIN", "http://localhost:3000"),
		},
	}

	return cfg, nil
}

func (d *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.Name,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("required environment variable %q is not set", key))
	}
	return v
}
