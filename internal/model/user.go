package model

import (
	"context"
	"quanta/internal/single"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            string    `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	Username      string    `json:"username"`
	FullName      string    `json:"fullName"`
	PreferredName string    `json:"preferredName"`
	Email         string    `json:"email"`
	Password      *string   `json:"password,omitempty"`
	Image         *string   `json:"image,omitempty"`
	CustomerId    *string
}

func GetUser(ctx context.Context, userId string) (User, error) {
	var user User

	query := `select id, created_at, username, full_name, preferred_name, email, password, image from users where id = $1;`
	row := single.DB.QueryRow(ctx, query, userId)
	err := row.Scan(&user.Id, &user.CreatedAt, &user.Username, &user.FullName, &user.PreferredName, &user.Email, &user.Password, &user.Image)
	if err != nil {
		return user, err
	}

	return user, nil
}

func VerifyUserAndCreateSession(ctx context.Context, email string, password string) (string, bool, error) {
	userId, verified, err := VerifyUser(ctx, email, password)
	if err != nil || !verified {
		return "", verified, err
	}

	id, err := CreateSession(ctx, userId)
	return id, verified, err
}

func CreateUserAndSession(ctx context.Context, user User) (User, string, error) {
	user, err := CreateUser(ctx, user)
	if err != nil {
		return User{}, "", err
	}

	id, err := CreateSession(ctx, user.Id)
	return user, id, err
}

func CreateUser(ctx context.Context, user User) (User, error) {
	var (
		id        string
		createdAt time.Time
	)

	hash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	query := `insert into users (username, full_name, preferred_name, email, password) values ($1, $2, $3, $4, $5) returning id, created_at`
	row := single.DB.QueryRow(ctx, query, user.Username, user.FullName, user.PreferredName, user.Email, hash)
	if err := row.Scan(&id, &createdAt); err != nil {
		return User{}, err
	}

	user.Id = id
	user.CreatedAt = createdAt
	return user, nil
}

func VerifyUser(ctx context.Context, email string, password string) (string, bool, error) {
	var (
		id   string
		hash string
	)

	query := `select id, password from users where email = $1;`
	err := single.DB.QueryRow(ctx, query, email).Scan(&id, &hash)
	if err == pgx.ErrNoRows {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return id, err == nil, nil
}
