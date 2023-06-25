package commomcase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/elliotchance/pie/v2"
)

type IGetUserList interface {
	Handle(accessToken string, roles []string) ([]*dtos.UserRes, error)
}

type GetUserList struct {
	authorService services.IAuthorizeService
	userRepo      repos.IUserRepo
}

func NewGetUserList(authorService services.IAuthorizeService, userRepo repos.IUserRepo) *GetUserList {
	return &GetUserList{authorService: authorService, userRepo: userRepo}
}

func (g GetUserList) Handle(accessToken string, roles []string) ([]*dtos.UserRes, error) {

	claims, err := g.authorService.Authorize(accessToken, []string{constants.RoleAdmin, constants.RolePlanner}, nil)
	if err != nil {
		return nil, err
	}

	if claims.Role == constants.RolePlanner && (pie.Contains(roles, constants.RoleAdmin) || pie.Contains(roles, constants.RolePlanner)) {
		return nil, errs.ErrUnauthorized
	}

	userList, _ := g.userRepo.FindAllByRoles(roles, nil)

	userResList := make([]*dtos.UserRes, len(userList))

	for i, user := range userList {
		userResList[i], _ = maps.ToUserRes(user)
	}

	return userResList, nil
}
