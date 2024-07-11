package email

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

const (
	changeLink = "http://localhost:80"
	mainPage   = "https://jiohealth.com/"
)

func SendMail(receiverMail string) {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GG_PASS := os.Getenv("GG_PASS")

	var body bytes.Buffer
	t, err := template.ParseFiles("./templates/responseEmail.html")

	if err != nil {
		panic(err)
	}

	t.Execute(&body, struct {
		Link     string
		MainPage string
	}{Link: changeLink, MainPage: mainPage})

	auth := ConfigEmail(GG_PASS)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: Reset your password" + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"coderbillzay@gmail.com",
		[]string{receiverMail},
		[]byte(msg),
	)

	if err != nil {
		panic(err)
	}
}
