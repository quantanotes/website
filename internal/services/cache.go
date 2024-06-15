package services

import (
	"quanta/internal/globals"

	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func init() {
	opts, err := redis.ParseURL(globals.Env("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	Cache = redis.NewClient(opts)
}
