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
	IsActive     bool
}

func NewUser(
	role string,
	email string,
	name string,
	passwordHash string,
	isActive bool,
	createdAt time.Time) *User {

	return &User{
		EntityBase:   NewEntityBase(createdAt),
		Role:         role,
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
		IsActive:     isActive}
}
