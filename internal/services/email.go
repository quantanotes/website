package services

import (
	"fmt"
	"net/smtp"
	"quanta/internal/globals"
	"quanta/internal/single"
)

func SendMail(to string, subject string, msg string) error {
	fullMsg := fmt.Sprintf("From: robots@quanta.uno\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s", to, subject, msg)

	return smtp.SendMail(
		globals.EmailHost+":"+globals.EmailPort,
		single.Email,
		"robots@quanta.uno",
		[]string{to},
		[]byte(fullMsg))
}
