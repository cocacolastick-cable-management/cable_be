package entities

import (
	"errors"
	"time"
)

const (
	RoleAdmin      = "admin"
	RolePlanner    = "planner"
	RoleSupplier   = "supplier"
	RoleContractor = "contractor"
)


var (
	ErrInvalidRole     = errors.New("invalid role")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")

	ErrDupEmail = errors.New("duplicated email")
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
