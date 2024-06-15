package services

import (
	"fmt"
	"net/smtp"
	"quanta/internal/globals"
)

var (
	Email     smtp.Auth
	emailAddr string
)

func SendMail(to string, subject string, msg string) error {
	fullMsg := fmt.Sprintf("From: robots@quanta.uno\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s", to, subject, msg)

	return smtp.SendMail(
		emailAddr,
		Email,
		"robots@quanta.uno",
		[]string{to},
		[]byte(fullMsg))
}

func init() {
	Email = smtp.PlainAuth("", globals.Env("EMAIL_USERNAME"), globals.Env("EMAIL_PASSWORD"), globals.Env("EMAIL_HOST"))
	emailAddr = globals.Env("EMAIL_HOST") + globals.Env("EMAIL_PORT")
}
