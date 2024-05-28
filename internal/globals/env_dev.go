//go:build dev

package globals

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	AppURL           string
	SSRURL           string
	MaxwellURL       string
	RedisURL         string
	PostgresURL      string
	EmailUsername    string
	EmailPassword    string
	EmailHost        string
	EmailPort        string
	StripeKey        string
	StripeWebhookKey string
)

func init() {
	godotenv.Load(".env.dev")
	AppURL = os.Getenv("APP_URL")
	SSRURL = os.Getenv("SSR_URL")
	MaxwellURL = os.Getenv("MAXWELL_URL")
	RedisURL = os.Getenv("REDIS_URL")
	PostgresURL = os.Getenv("POSTGRES_URL")
	EmailUsername = os.Getenv("EMAIL_USERNAME")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
	EmailHost = os.Getenv("EMAIL_HOST")
	EmailPort = os.Getenv("EMAIL_PORT")
	StripeKey = os.Getenv("STRIPE_KEY")
	StripeWebhookKey = os.Getenv("STRIPE_WEBHOOK_KEY")
}
