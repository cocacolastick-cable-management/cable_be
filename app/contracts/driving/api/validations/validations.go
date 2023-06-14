package validations

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/elliotchance/pie/v2"
	"github.com/go-playground/validator/v10"
)

type VlCreateUserDepd struct {
	userRepo repos.IUserRepo
}

func NewVlCreateUserDepd(userRepo repos.IUserRepo) *VlCreateUserDepd {
	return &VlCreateUserDepd{userRepo: userRepo}
}

func (v VlCreateUserDepd) Handle(sl validator.StructLevel) {

	req := sl.Current().Interface().(dtos.CreateUserReq)

	matchUser, _ := v.userRepo.FindByEmail(req.Email, nil)
	if matchUser != nil {
		sl.ReportError(req.Email, "email", "Email", "duplicated", "duplicated email")
	}

	if !pie.Contains(constants.RoleList, req.Role) {
		sl.ReportError(req.Role, "role", "Role", "invalid", "invalid role")
	}

	// TODO should I validate dtos.CreateUserReq.Name?
}
