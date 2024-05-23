package single

import (
	"net/smtp"
	"quanta/internal/globals"
)

var (
	Email smtp.Auth
)

func init() {
	Email = smtp.PlainAuth("", globals.EmailUsername, globals.EmailPassword, globals.EmailHost)
}
