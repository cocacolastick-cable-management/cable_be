package services

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/elliotchance/pie/v2"
	"github.com/go-playground/validator/v10"
	"time"
)

type IUserFactory interface {
	CreateUser(role, email, name, password string) (user *entities.User, err error)
}

type UserFactory struct {
	passwordService IPasswordService
	userRepo        repos.IUserRepo
	validator       *validator.Validate
}

func NewUserFactory(
	passwordService IPasswordService,
	userRepo repos.IUserRepo,
	validator *validator.Validate) *UserFactory {

	return &UserFactory{
		passwordService: passwordService,
		userRepo:        userRepo,
		validator:       validator}
}

func (u UserFactory) CreateUser(role, email, name, password string) (user *entities.User, err error) {

	err = u.validator.Var(email, "required,email")
	if err != nil {
		return nil, errs.ErrInvalidEmail
	}

	matchUser, _ := u.userRepo.FindByEmail(email, nil)
	if matchUser != nil {
		return nil, errs.ErrDupEmail
	}

	if !pie.Contains(constants.RoleList, role) {
		return nil, errs.ErrInvalidRole
	}

	// TODO: should I validate name?

	passwordHash, err := u.passwordService.Hash(password)
	if err != nil {
		return nil, err
	}

	user = entities.NewUser(role, email, name, passwordHash, true, time.Now())

	return user, nil
}
