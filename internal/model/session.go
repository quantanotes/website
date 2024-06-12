package model

import (
	"context"
	"quanta/internal/services"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const sessionPrefix = "session-"

func CreateSession(ctx context.Context, user string) (string, error) {
	id := uuid.New().String()
	err := SetSession(ctx, id, user)
	return id, err
}

func SetSession(ctx context.Context, id string, user string) error {
	key := sessionPrefix + id
	expires := time.Hour * 24 * 30
	return services.Cache.Set(ctx, key, user, expires).Err()
}

func GetSession(ctx context.Context, id string) (string, bool, error) {
	key := sessionPrefix + id
	user, err := services.Cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	}
	return user, true, nil
}

func DeleteSession(ctx context.Context, id string) error {
	key := sessionPrefix + id
	return services.Cache.Del(ctx, key).Err()
}
