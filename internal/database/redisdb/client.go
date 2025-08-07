package redisdb

import (
	"context"
	"log"
	"time"

	"github.com/Amir-Sadati/order-packing/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Printf("Redis connection failed: %v", err)
		return nil, err
	}

	return rdb, nil
}
