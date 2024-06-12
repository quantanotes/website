package model

import (
	"reflect"
	"strings"

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

func columns(data any) string {
	return strings.Join(fields(data), ", ")
}

func fields(data any) []string {
	var fields []string
	typ := reflect.TypeOf(data)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		col := field.Tag.Get("db")
		if col == "" {
			col = snakecase(field.Name)
		}
		fields = append(fields, snakecase(col))
	}
	return fields
}

func snakecase(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			builder.WriteRune('_')
		}
		builder.WriteRune(r)
	}
	return strings.ToLower(builder.String())
}
