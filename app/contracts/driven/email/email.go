package email

import "github.com/cable_management/cable_be/app/domain/entities"

type MailData struct {
	Receiver string
	Subject  string
	Body     string
}

type IEmail interface {
	//send(data *MailData) error
	SendEmailNewUser(emailDto MailNewUserDto) error
	SendEmailUpdateUserIsActive(emailDto MailUpdateUserIsActiveDto) error
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
