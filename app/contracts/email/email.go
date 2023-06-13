package email

import "github.com/cable_management/cable_be/app/domain/entities"

type EmailData struct {
	Receiver string
	Subject  string
	Body     string
}

type IEmail interface {
	//send(data *EmailData) error
	SendEmailNewUser(emailDto EmailNewUserDto) error
}

type EmailNewUserDto struct {
	Name     string
	Email    string
	Password string
}

func ToEmailNewUserDto(user *entities.User, password string) EmailNewUserDto {
	return EmailNewUserDto{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}
}
