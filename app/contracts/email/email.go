package email

type EmailData struct {
	Receiver string
	Subject  string
	Body     string
}

type IEmail interface {
	Send(data *EmailData) error
}
