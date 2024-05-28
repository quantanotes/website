package model

import (
	"context"
	"quanta/internal/single"

	"github.com/stripe/stripe-go/v78"
)

func GetOrCreateCustomerId(ctx context.Context, userId string) (string, error) {
	id, err := GetCustomerId(ctx, userId)
	if err != nil {
		return "", err
	} else if id == "" {
		return CreateCustomer(ctx, userId)
	}
	return id, nil
}

func CreateCustomer(ctx context.Context, userId string) (string, error) {
	user, err := GetUser(ctx, userId)
	if err != nil {
		return "", err
	}

	params := &stripe.CustomerParams{
		Name:  stripe.String(user.FullName),
		Email: stripe.String(user.Email),
	}

	cus, err := single.Stripe.Customers.New(params)
	if err != nil {
		return "", err
	}

	return cus.ID, UpdateCustomerId(ctx, cus.ID, userId)
}

func GetCustomerId(ctx context.Context, userId string) (string, error) {
	query := `
		SELECT customer_id
		FROM users
		WHERE id = $1
	`
	var cusId string
	err := single.DB.QueryRow(ctx, query, userId).Scan(&cusId)
	return cusId, err
}

func UpdateCustomerId(ctx context.Context, cusId string, userId string) error {
	query := `
		UPDATE users
		SET customer_id = $1
		WHERE id = $2
	`
	_, err := single.DB.Exec(ctx, query, cusId, userId)
	return err
}
