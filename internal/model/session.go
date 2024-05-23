package model

import (
	"context"
	"quanta/internal/single"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const sessionPrefix = "session-"

func CreateSession(ctx context.Context, user string) (string, error) {
	id := uuid.New().String()
	return id, SetSession(ctx, id, user)
}

func SetSession(ctx context.Context, id string, user string) error {
	return single.Cache.Set(ctx, sessionPrefix+id, user, time.Hour*24*30).Err()
}

func GetSession(ctx context.Context, id string) (string, bool, error) {
	user, err := single.Cache.Get(ctx, sessionPrefix+id).Result()
	if err == redis.Nil {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	}
	return user, true, nil
}

func DeleteSession(ctx context.Context, id string) error {
	return single.Cache.Del(ctx, sessionPrefix+id).Err()
}
