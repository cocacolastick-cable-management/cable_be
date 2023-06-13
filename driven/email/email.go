package email

import (
	"github.com/cable_management/cable_be/app/contracts/email"
	"log"
	"net/smtp"
)

type EmailConfig struct {
	MailHost string
	Host     string
	Port     string
	Password string
}

type Email struct {
	config EmailConfig
	auth   smtp.Auth
}

func NewEmail(config EmailConfig) *Email {
	return &Email{
		config: config,
		auth:   smtp.PlainAuth("", config.MailHost, config.Password, config.Host)}
}

func (e Email) Send(data *email.EmailData) error {

	mail := "From: " + e.config.MailHost + "\n" +
		"To: " + data.Receiver + "\n" +
		"Subject: " + data.Subject + "\n" +
		"\n" +
		data.Body

	err := smtp.SendMail(e.config.Host+":"+e.config.Port, e.auth, e.config.MailHost, []string{data.Receiver}, []byte(mail))
	if err != nil {
		log.Fatal(err)
	}

	return err
}
