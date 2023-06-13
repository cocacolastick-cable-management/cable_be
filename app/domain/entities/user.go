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
	RoleList = []string{RoleAdmin, RolePlanner, RoleSupplier, RoleContractor}
)

var (
	ErrInvalidRole     = errors.New("invalid role")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")

	ErrDupEmail = errors.New("duplicated email")
)

type User struct {
	EntityBase

	Role         string `gorm:"type:varchar"`
	Email        string `gorm:"type:varchar,unique"`
	Name         string `gorm:"type:varchar"`
	PasswordHash string `gorm:"type:varchar"`
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
