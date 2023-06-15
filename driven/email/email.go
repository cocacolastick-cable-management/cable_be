package email

import (
	"fmt"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/domain/constants"
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

func (e Email) SendEmailOnRequestUpdate(mailDtoList []email.MailRequestActionDto) error {

	for _, mail := range mailDtoList {

		mail := mail
		go func() {

			body := ""
			switch mail.Action {
			case constants.ActionCreate:
				// body for ActionCreate
				body = fmt.Sprintf("%v just create a new request-%v on contract-%v", mail.SenderEmail, mail.RequestCounter, mail.ContractCounter)
			case constants.ActionUpdate:
				// body for ActionUpdate
				body = fmt.Sprintf("%v just mark request-%v as action%v", mail.SenderEmail, mail.RequestCounter, mail.Status)
			case constants.ActionCancel:
				// body for CancelAction
				body = fmt.Sprintf("%v cancel the request-%v on contract-%v", mail.SenderEmail, mail.RequestCounter, mail.ContractCounter)
			}

			_ = e.send(&email.MailData{
				Receiver: mail.ReceiverEmail,
				Subject:  mail.SenderEmail,
				Body:     body,
			})
		}()
	}

	return nil
}
