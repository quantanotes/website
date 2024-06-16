package model

import (
	"github.com/jackc/pgx/v5"
)

type scannable interface {
	scan() []any
}

func scanMany[T any](rows pgx.Rows) ([]T, error) {
	var res []T
	for rows.Next() {
		var r T
		if err := rows.Scan(any(&r).(scannable).scan()...); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
