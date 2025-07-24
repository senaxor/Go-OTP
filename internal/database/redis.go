package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	redisAddr := os.Getenv("REDIS_ADDR") // e.g., "redis:6379"
	if redisAddr == "" {
		redisAddr = "localhost:6380"
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Optional: test connection
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
