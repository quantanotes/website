package single

import (
	"context"
	"quanta/internal/globals"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func init() {
	var err error
	if DB, err = pgxpool.New(context.Background(), globals.PostgresURL); err != nil {
		panic(err)
	}
}
