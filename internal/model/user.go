package model

import (
	"context"
	"quanta/internal/services"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	userColumns = columns(User{})
)

type User struct {
	ID            string    `json:"id" db:"id"`
	Created       time.Time `json:"created" db:"created_at"`
	Username      string    `json:"username"`
	FullName      string    `json:"fullName"`
	PreferredName string    `json:"preferredName"`
	Email         string    `json:"email"`
	Password      *string   `json:"password,omitempty"`
	Image         *string   `json:"image,omitempty"`
	CustomerID    *string   `json:"customerID,omitempty" db:"customer_id"`
}

func GetUser(ctx context.Context, userID string) (User, error) {
	var user User
	query := `
		SELECT ` + userColumns + `
		FROM users
		WHERE id = $1
	`
	row := services.DB.QueryRow(ctx, query, userID)
	err := row.Scan(user.scan()...)
	return user, err
}

func CreateUser(ctx context.Context, user User) (User, error) {
	query := `
		INSERT INTO users
		(username, full_name, preferred_name, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	cols := []any{user.Username, user.FullName, user.PreferredName, user.Email, user.Password}
	row := services.DB.QueryRow(ctx, query, cols...)
	err := row.Scan(&user.ID)
	return user, err
}

func VerifyUser(ctx context.Context, email string, password string) (string, bool, error) {
	var id, hash string
	query := `SELECT id, password FROM users WHERE email = $1`
	err := services.DB.QueryRow(ctx, query, email).Scan(&id, &hash)
	if err == pgx.ErrNoRows {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return id, err == nil, nil
}

func VerifyUserAndCreateSession(ctx context.Context, email string, password string) (string, bool, error) {
	userID, verified, err := VerifyUser(ctx, email, password)
	if err != nil || !verified {
		return "", verified, err
	}
	id, err := CreateSession(ctx, userID)
	return id, verified, err
}

func CreateUserAndSession(ctx context.Context, user User) (User, string, error) {
	user, err := CreateUser(ctx, user)
	if err != nil {
		return User{}, "", err
	}
	id, err := CreateSession(ctx, user.ID)
	return user, id, err
}

func (u *User) scan() []any {
	return []any{
		&u.ID,
		&u.Created,
		&u.Username,
		&u.FullName,
		&u.PreferredName,
		&u.Email,
		&u.Password,
		&u.Image,
		&u.CustomerID,
	}
}
