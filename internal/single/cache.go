package single

import (
	"quanta/internal/globals"

	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func init() {
	opts, err := redis.ParseURL(globals.RedisURL)
	if err != nil {
		panic(err)
	}
	Cache = redis.NewClient(opts)
}
