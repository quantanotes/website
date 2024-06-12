//go:build dev

package globals

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	IoURL string
)

func Env(key string) string {
	return os.Getenv(key)
}

func init() {
	godotenv.Load(".env.dev")
	IoURL = os.Getenv("IO_URL")
}
