package services

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/elliotchance/pie/v2"
)

type IAuthorizeService interface {
	Authorize(accessToken string, targetRoles []string, targetPerms []string) (claims *AuthTokenClaims, err error)
}

type AuthorizeService struct {
	tokenService IAuthTokenService
	userRepo     repos.IUserRepo
}

func NewAuthorizeService(tokenService IAuthTokenService, userRepo repos.IUserRepo) *AuthorizeService {
	return &AuthorizeService{tokenService: tokenService, userRepo: userRepo}
}

func (a AuthorizeService) Authorize(accessToken string, targetRoles []string, targetPerms []string) (claims *AuthTokenClaims, err error) {

	isValid, claims := a.tokenService.IsAccessTokenValid(accessToken)
	if !isValid {
		return nil, errs.ErrUnauthenticated
	}

	if (targetRoles != nil && !pie.Contains(targetRoles, claims.Role)) ||
		(targetPerms != nil && !pie.Equals(targetPerms, claims.Permissions)) {
		return nil, errs.ErrUnauthorized
	}

	matchUser, _ := a.userRepo.FindById(claims.UserId, nil)

	if matchUser == nil {
		return nil, errs.ErrUserNotFound
	}

	if !matchUser.IsActive {
		return nil, errs.ErrUserIsDisable
	}

	return claims, nil
}
