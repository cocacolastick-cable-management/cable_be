package services

import (
	"crypto/rand"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

const (
	PasswordValidateTags = "required"
)

type IPasswordService interface {
	Hash(password string) (passwordHash string, err error)
	Compare(password, hashPassword string) bool
	GeneratePassword(length int) string
}

type PasswordService struct {
	validator *validator.Validate
}

func NewPasswordHash(validator *validator.Validate) *PasswordService {
	return &PasswordService{validator: validator}
}

func (p PasswordService) Hash(password string) (passwordHash string, err error) {

	err = p.validator.Var(password, PasswordValidateTags)
	if err != nil {
		return "", errs.ErrInvalidPassword
	}

	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	passwordHash = string(salt)
	return passwordHash, nil
}

func (p PasswordService) Compare(password, hashPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func (p PasswordService) GeneratePassword(length int) string {

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password)
}
