package admincase

import (
	"github.com/cable_management/cable_be/app/contracts/database/repos"
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
	userRepo      repos.IUserRepo
	userFac       services.IUserFactory
	validator     *validator.Validate
	authorService services.IAuthorizeService
}

func NewCreateUser(
	userRepo repos.IUserRepo,
	userFac services.IUserFactory,
	validator *validator.Validate,
	authorService services.IAuthorizeService) *CreateUser {

	return &CreateUser{
		userRepo:      userRepo,
		userFac:       userFac,
		validator:     validator,
		authorService: authorService}
}

func (c CreateUser) Handle(accessToken string, req *CreateUserReq) (err error) {

	_, err = c.authorService.Authorize(accessToken, nil, nil)
	if err != nil {
		return err
	}

	err = c.validator.Struct(req)
	if err != nil {
		return err
	}

	newUser, err := c.userFac.CreateUser(req.Role, req.Email, req.Name, "")
	if err != nil {
		return err
	}

	err = c.userRepo.Insert(newUser)
	// TODO send email with email and password to user
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
