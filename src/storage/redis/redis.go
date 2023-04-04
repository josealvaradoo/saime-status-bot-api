package redis

import (
	"fmt"

	"github.com/josealvaradoo/saime-status-bot/src/utils"
	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
)

type Driver string

const (
	_redis Driver = "redis"
)

func New(driver Driver) {
	switch driver {
	case _redis:
		newRedisConnection()
	}
}

func newRedisConnection() {
	host := utils.Env(utils.REDISHOST)
	port := utils.Env(utils.REDISPORT)
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
	})
}
