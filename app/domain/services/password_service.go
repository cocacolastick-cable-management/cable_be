package services

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

const (
	PasswordValidateTags = ""
)

type IPasswordService interface {
	Hash(password string) (passwordHash string, err error)
	Compare(password, hashPassword string) bool
}

type PasswordHash struct {
	validator *validator.Validate
}

func NewPasswordHash(validator *validator.Validate) *PasswordHash {
	return &PasswordHash{validator: validator}
}

func (p PasswordHash) Hash(password string) (passwordHash string, err error) {

	err = p.validator.Var(password, PasswordValidateTags)
	if err != nil {
		return "", entities.ErrInvalidPassword
	}

	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	passwordHash = string(salt)
	return passwordHash, nil
}

func (p PasswordHash) Compare(password, hashPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
