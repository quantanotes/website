package model

import (
	"context"
	"encoding/json"
	"fmt"
	"quanta/internal/services"
	"quanta/internal/single"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const verificationPrefix = "verification-"

func SendVerificationEmail(ctx context.Context, user User) error {
	verificationCode, err := CreateVerificationCode(ctx, user)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(
		`Here's your verification code: %s. This code will expire in 30 minutes. Alternatively, you can click <a href="https://www.quanta.uno/verify/%s">here</a> to verify.`,
		verificationCode,
		verificationCode)

	return services.SendMail(user.Email, "Your Quanta Email Verification Code!", msg)
}

func GetVerificationUser(ctx context.Context, id string) (User, bool, error) {
	var user User

	res, err := single.Cache.Get(ctx, verificationPrefix+id).Result()
	if err == redis.Nil {
		return user, false, nil
	} else if err != nil {
		return user, true, err
	}

	err = json.Unmarshal([]byte(res), &user)
	return user, true, err
}

func CreateVerificationCode(ctx context.Context, user User) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	id := uuid.New().String()
	err = single.Cache.Set(ctx, verificationPrefix+id, data, time.Minute*30).Err()
	return id, err
}
