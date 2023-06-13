package admincase

import (
	"fmt"
	"github.com/cable_management/cable_be/app/contracts/database/repos"
	"github.com/cable_management/cable_be/app/contracts/email"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/elliotchance/pie/v2"
	"github.com/go-playground/validator/v10"
)

type CreateUserReq struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ICreateUser interface {
	Handle(accessToken string, req *CreateUserReq) (err error)
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

func (c CreateUser) Handle(accessToken string, req *CreateUserReq) (err error) {

	// authorize
	_, err = c.authorService.Authorize(accessToken, []string{entities.RoleAdmin}, nil)
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
		err := c.emailDriven.SendEmailNewUser(email.ToEmailNewUserDto(newUser, password))
		if err != nil {
			// TODO logger error
			fmt.Println(err)
		}
	}()
	// TODO sse?

	return nil
}

// validation

type VlCreateUserDepd struct {
	userRepo repos.IUserRepo
}

func NewVlCreateUserDepd(userRepo repos.IUserRepo) *VlCreateUserDepd {
	return &VlCreateUserDepd{userRepo: userRepo}
}

func (v VlCreateUserDepd) Handle(sl validator.StructLevel) {

	req := sl.Current().Interface().(CreateUserReq)

	matchUser, _ := v.userRepo.FindByEmail(req.Email, nil)
	if matchUser != nil {
		sl.ReportError(req.Email, "email", "Email", "duplicated", "duplicated email")
	}

	if !pie.Contains(entities.RoleList, req.Role) {
		sl.ReportError(req.Role, "role", "Role", "invalid", "invalid role")
	}

	// TODO should I validate CreateUserReq.Name?
}
