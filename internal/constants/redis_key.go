// Package constants provides application-wide constants
package constants

// RedisKey represents a Redis key type
type RedisKey string

const (
	// RedisKeyPackSizes is the Redis key for storing pack sizes
	RedisKeyPackSizes RedisKey = "pack_sizes"
)
