package email

import (
	"fmt"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"log"
	"net/smtp"
)

type Config struct {
	MailHost string
	Host     string
	Port     string
	Password string
}

type Email struct {
	config Config
	auth   smtp.Auth
}

func NewEmail(config Config) *Email {
	return &Email{
		config: config,
		auth:   smtp.PlainAuth("", config.MailHost, config.Password, config.Host)}
}

func (e Email) send(data *email.MailData) error {

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

func (e Email) SendEmailNewUser(emailDto email.MailNewUserDto) error {

	err := e.send(&email.MailData{
		Receiver: emailDto.Email,
		Subject:  "Your Account",
		Body:     fmt.Sprintf("\n name: %v \n email: %v \n password: %v\n", emailDto.Name, emailDto.Email, emailDto.Password),
	})

	return err
}

func (e Email) SendEmailUpdateUserIsActive(emailDto email.MailUpdateUserIsActiveDto) error {

	status := "disable"
	if emailDto.NewStatus {
		status = "active"
	}

	err := e.send(&email.MailData{
		Receiver: emailDto.Email,
		Subject:  fmt.Sprintf("your account is %v", status),
		Body:     fmt.Sprintf("your account is %v", status),
	})

	return err
}