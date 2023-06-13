package email

type EmailNewUserDto struct {
	Name     string
	Email    string
	Password string
}

type EmailData struct {
	Receiver string
	Subject  string
	Body     string
}

type IEmail interface {
	//send(data *EmailData) error
	SendEmailNewUser(emailDto EmailNewUserDto) error
}
