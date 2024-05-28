//go:build !dev

package globals

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	AppURL              string
	SSRURL              string
	MaxwellURL          string
	RedisURL            string
	PostgresURL         string
	EmailUsername       string
	EmailPassword       string
	EmailHost           string
	EmailPort           string
	StripeSecretKey     string
	StripeWebhookSecret string
)

func init() {
	godotenv.Load()
	AppURL = os.Getenv("APP_URL")
	SSRURL = os.Getenv("SSR_URL")
	MaxwellURL = os.Getenv("MAXWELL_URL")
	RedisURL = os.Getenv("REDIS_URL")
	PostgresURL = os.Getenv("POSTGRES_URL")
	EmailUsername = os.Getenv("EMAIL_USERNAME")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
	EmailHost = os.Getenv("EMAIL_HOST")
	EmailPort = os.Getenv("EMAIL_PORT")
	StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
	StripeWebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
}
