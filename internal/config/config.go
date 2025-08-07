package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP  *HttpConfig
	Redis *RedisConfig
}

type HttpConfig struct {
	Host string
	Port string
}

type RedisConfig struct {
	Address  string
	Password string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		HTTP:  loadHttpConfig(),
		Redis: loadRedisConfig(),
	}, nil
}

func loadHttpConfig() *HttpConfig {
	return &HttpConfig{
		Host: getEnv("HTTP_HOST"),
		Port: getEnv("HTTP_PORT"),
	}
}

func loadRedisConfig() *RedisConfig {
	return &RedisConfig{
		Address:  getEnv("REDIS_ADDRESS"),
		Password: getEnv("REDIS_PASSWORD"),
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
