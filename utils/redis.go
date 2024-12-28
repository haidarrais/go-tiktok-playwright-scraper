package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// RedisClient is the Redis client instance
var RedisClient *redis.Client

// InitializeRedis initializes the Redis client
func InitializeRedis(addr string, password string, db int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
}

// SetCache sets a value in Redis with an expiration time
func SetCache(key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// GetCache retrieves a value from Redis
func GetCache(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	}
	return val, err
}
