package services

import (
	"context"
	"quanta/internal/globals"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func init() {
	var err error
	ctx := context.Background()
	url := globals.Env("POSTGRES_URL")
	if DB, err = pgxpool.New(ctx, url); err != nil {
		panic(err)
	}
}
