// Package redisdb provides Redis database connectivity and operations
package redisdb

import (
	"context"
	"log"
	"time"

	"github.com/Amir-Sadati/order-packing/internal/config"
	"github.com/redis/go-redis/v9"
)

// NewClient creates and returns a new Redis client with the given configuration
func NewClient(ctx context.Context, cfg *config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Printf("Redis connection failed: %v", err)
		return nil, err
	}

	return rdb, nil
}
