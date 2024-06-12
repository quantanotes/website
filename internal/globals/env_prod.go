//go:build !dev

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
	godotenv.Load()
	IoURL = os.Getenv("IO_URL")
}
