package cache

import (
	"fmt"

	"github.com/josealvaradoo/saime-status-bot/src/utils"
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
	host := utils.Env(utils.REDIS_HOST)
	port := utils.Env(utils.REDIS_PORT)
	cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Cache() *redis.Client {
	return cache
}
