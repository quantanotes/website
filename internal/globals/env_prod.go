//go:build !dev

package globals

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	SSRURL        string
	MaxwellURL    string
	RedisURI      string
	PostgresURL   string
	EmailUsername string
	EmailPassword string
	EmailHost     string
	EmailPort     string
)

func init() {
	godotenv.Load()
	SSRURL = os.Getenv("SSR_URL")
	MaxwellURL = os.Getenv("MAXWELL_URL")
	RedisURI = os.Getenv("REDIS_URL")
	PostgresURL = os.Getenv("POSTGRES_URL")
	EmailUsername = os.Getenv("EMAIL_USERNAME")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
	EmailHost = os.Getenv("EMAIL_HOST")
	EmailPort = os.Getenv("EMAIL_PORT")
}
