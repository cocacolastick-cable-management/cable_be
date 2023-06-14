package admincase

import (
	"fmt"
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/go-playground/validator/v10"
)

type ICreateUser interface {
	Handle(accessToken string, req *dtos.CreateUserReq) (err error)
}

type CreateUser struct {
	userRepo        repos.IUserRepo
	userFac         services.IUserFactory
	validator       *validator.Validate
	authorService   services.IAuthorizeService
	emailDriven     email.IEmail
	passwordService services.IPasswordService
}

func NewCreateUser(
	userRepo repos.IUserRepo,
	userFac services.IUserFactory,
	validator *validator.Validate,
	authorService services.IAuthorizeService,
	emailDriven email.IEmail,
	passwordService services.IPasswordService) *CreateUser {

	return &CreateUser{
		userRepo:        userRepo,
		userFac:         userFac,
		validator:       validator,
		authorService:   authorService,
		emailDriven:     emailDriven,
		passwordService: passwordService}
}

func (c CreateUser) Handle(accessToken string, req *dtos.CreateUserReq) (err error) {

	// authorize
	_, err = c.authorService.Authorize(accessToken, []string{constants.RoleAdmin}, nil)
	if err != nil {
		return err
	}

	// validate
	err = c.validator.Struct(req)
	if err != nil {
		return err
	}

	// create user
	password := c.passwordService.GeneratePassword(10)
	newUser, err := c.userFac.CreateUser(req.Role, req.Email, req.Name, password)
	if err != nil {
		return err
	}

	// insert to database
	err = c.userRepo.Insert(newUser)
	if err != nil {
		return err
	}

	// send email with account to user
	go func() {
		err := c.emailDriven.SendEmailNewUser(email.ToMailNewUserDto(newUser, password))
		if err != nil {
			// TODO logger error
			fmt.Println(err)
		}
	}()
	// TODO sse?

	return nil
}

// validation
