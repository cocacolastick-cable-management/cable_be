package entities


type User struct {
	EntityBase

	Role         string
	Email        string
	Name         string
	PasswordHash string
}
