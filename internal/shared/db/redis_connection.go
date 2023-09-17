package db_connection

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/io-m/hyphen/pkg/constants"
)

func CreateRedisConnection() (*redis.Client, error) {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(constants.REDIS_ADDRESS),  // Redis server address
		Password: os.Getenv(constants.REDIS_PASSWORD), // No password set
		DB:       0,                                   // Default DB
	})

	// Ping the Redis server to test the connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("error pinging Redis: %w", err)
	}

	return client, nil
}
