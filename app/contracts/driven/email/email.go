package email

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"time"
)

type MailData struct {
	Receiver string
	Subject  string
	Body     string
}

type IEmail interface {
	SendEmailNewUser(emailDto MailNewUserDto) error
	SendEmailUpdateUserIsActive(emailDto MailUpdateUserIsActiveDto) error
	SendEmailOnRequestUpdate(mailDtoList []MailRequestActionDto) error
}

type MailNewUserDto struct {
	Name     string
	Email    string
	Password string
}

func ToMailNewUserDto(user *entities.User, password string) MailNewUserDto {
	return MailNewUserDto{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}
}

type MailUpdateUserIsActiveDto struct {
	NewStatus bool
	Email     string
}

///

type MailRequestActionDto struct {
	SenderEmail     string
	ReceiverEmail   string
	Action          string
	Status          string
	RequestCounter  uint
	ContractCounter uint
	Time            time.Time
}

///
