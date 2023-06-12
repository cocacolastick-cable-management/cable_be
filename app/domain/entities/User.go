package entities

import "time"

const (
	RoleAdmin      = "admin"
	RolePlanner    = "planner"
	RoleSupplier   = "supplier"
	RoleContractor = "contractor"
)

type User struct {
	EntityBase

	Role         string
	Email        string
	Name         string
	PasswordHash string
}

func NewUser(
	role,
	email,
	name,
	passwordHash string,
	createdAt time.Time) *User {

	return &User{
		NewEntityBase(createdAt),
		role,
		email,
		name,
		passwordHash}
}
