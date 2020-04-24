package emails

import (
	log "github.com/sirupsen/logrus"
	"net/smtp"
)

type Box struct {
}

func (b *Box) SendMail(body string) {
	from := "bettercallssx@gmail.com"
	to := "bettercallssx@163.com"
	pass := "..."
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
