package email

import (
	"net/smtp"
)

func ConfigEmail(gmailPass string) smtp.Auth {
	auth := smtp.PlainAuth(
		"",
		"coderbillzay@gmail.com",
		gmailPass,
		"smtp.gmail.com",
	)

	return auth
}
