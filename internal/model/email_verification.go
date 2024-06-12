package model

import (
	"context"
	"encoding/json"
	"fmt"
	"quanta/internal/services"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

//TODO: decouple verification id and verification code

const emailVerificationPrefix = "email-verification-"

func SendEmailVerification(ctx context.Context, user User, redirect string) error {
	code, err := CreateEmailVerificationCode(ctx, user)
	if err != nil {
		return err
	}
	link := "https://www.quanta.uno/verify/" + code
	if redirect != "" {
		link += "?redirect=" + redirect
	}
	msg := fmt.Sprintf(
		`<a href="%s">Here's</a> your verification link. This link expires in 30 minutes. If you did not request this email, please ignore it.`,
		link,
	)
	return services.SendMail(user.Email, "Your Quanta Email Verification Code!", msg)
}

func GetUserEmailVerificationCode(ctx context.Context, id string) (User, bool, error) {
	var user User
	res, err := services.Cache.Get(ctx, emailVerificationPrefix+id).Result()
	if err == redis.Nil {
		return user, false, nil
	} else if err != nil {
		return user, true, err
	}
	err = json.Unmarshal([]byte(res), &user)
	return user, true, err
}

func CreateEmailVerificationCode(ctx context.Context, user User) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	id := uuid.New().String()
	err = services.Cache.Set(ctx, emailVerificationPrefix+id, data, time.Minute*30).Err()
	return id, err
}
