package redis

import (
	"github.com/go-redis/redis"
	"os"
)

// Client defines the Redis client.
var Client *redis.Client

// Re-defines nil so we do not have to import Redis multiple times.
var Nil = redis.Nil

func init() {
	Addr := os.Getenv("REDIS_ADDR")
	if Addr == "" {
		panic("REDIS_ADDR is nil.")
	}
	Password := os.Getenv("REDIS_PASSWORD")
	Client = redis.NewClient(&redis.Options{
		Addr: Addr,
		Password: Password,
		DB: 0,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
