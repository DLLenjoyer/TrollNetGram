package config

import (
    "os"
)

type Config struct {
    Port     string
    DBHost   string
    DBPort   string
    DBUser   string
    DBName   string
    DBPass   string
    SSLMode  string
}

func NewConfig() *Config {
    return &Config{
        Port:     getEnv("PORT", "8080"),
        DBHost:   getEnv("DB_HOST", "localhost"),
        DBPort:   getEnv("DB_PORT", "5432"),
        DBUser:   getEnv("DB_USER", "user"),
        DBName:   getEnv("DB_NAME", "mydb"),
        DBPass:   getEnv("DB_PASS", ""),
        SSLMode:  getEnv("SSL_MODE", "disable"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
