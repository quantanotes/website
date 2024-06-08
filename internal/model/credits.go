package model

import (
	"context"
	"quanta/internal/single"
)

func AddCredits(customerId string, amount int) error {
	query := `
		UPDATE users
		SET credits = credits + $1
		WHERE customer_id = $2
	`
	_, err := single.DB.Exec(context.Background(), query)
	return err
}

func UseCredits(userId string, amount int) error {
	query := `
		UPDATE users
		SET credits = credits - $1
		WHERE customer_id = $2
	`
	_, err := single.DB.Exec(context.Background(), query)
	return err
}
