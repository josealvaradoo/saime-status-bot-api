package storage

import (
	"github.com/redis/go-redis/v9"
)

var (
	cache *redis.Client
)

type Driver string

const (
	Redis Driver = "REDIS"
)

func New(driver Driver) {
	switch driver {
	case Redis:
		newRedisConnection()
	}
}

func newRedisConnection() {
	cache = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Cache() *redis.Client {
	return cache
}
