// Package config provides configuration management for the application
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the main application configuration
type Config struct {
	HTTP  *HTTPConfig
	Redis *RedisConfig
}

// HTTPConfig represents HTTP server configuration
type HTTPConfig struct {
	Host string
	Port string
}

// RedisConfig represents Redis connection configuration
type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		HTTP:  loadHTTPConfig(),
		Redis: loadRedisConfig(),
	}, nil
}

func loadHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		Host: getEnv("HTTP_HOST"),
		Port: getEnv("HTTP_PORT"),
	}
}

func loadRedisConfig() *RedisConfig {
	return &RedisConfig{
		Address:  getEnv("REDIS_ADDRESS"),
		Password: getEnv("REDIS_PASSWORD"),
		DB:       getEnvAsInt("REDIS_DB"),
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing env var: %s", key)
	}

	return val
}

func getEnvAsInt(key string) int {
	val := getEnv(key)
	n, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Invalid int for %s: %v", key, err)
	}
	return n
}
