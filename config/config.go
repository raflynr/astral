package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Postgres struct {
		Host     string
		User     string
		Password string
		Port     string
		Name     string
	}
	Fiber struct {
		Port string
	}
	JWT struct {
		Secret string
	}
}

func NewConfig() AppConfig {
	_ = godotenv.Load(".env")

	return AppConfig{
		Postgres: struct {
			Host     string
			User     string
			Password string
			Port     string
			Name     string
		}{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		Fiber: struct {
			Port string
		}{
			Port: os.Getenv("PORT"),
		},
		JWT: struct {
			Secret string
		}{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
