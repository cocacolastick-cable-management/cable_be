package entities

import (
	"time"
)

type User struct {
	EntityBase

	Role         string `gorm:"type:varchar"`
	Email        string `gorm:"type:varchar;unique"`
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
