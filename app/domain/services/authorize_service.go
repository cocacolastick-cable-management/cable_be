package services

import (
	"errors"
	"github.com/cable_management/cable_be/app/contracts/database/repos"
	"github.com/elliotchance/pie/v2"
)

type IAuthorizeService interface {
	Authorize(accessToken string, targetRoles []string, targetPerms []string) (claims *AuthTokenClaims, err error)
}

var (
	ErrUnauthenticated = errors.New("authenticate failed")
	ErrUnauthorized    = errors.New("authenticate failed")
	ErrUserIsDisable   = errors.New("authenticate failed")
)

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
		return nil, ErrUnauthenticated
	}

	if (targetRoles != nil && !pie.Contains(targetRoles, claims.Role)) ||
		(targetPerms != nil && !pie.Equals(targetPerms, claims.Permissions)) {
		return nil, ErrUnauthorized
	}

	matchUser, _ := a.userRepo.FindById(claims.UserId, nil)
	if !matchUser.IsActive {
		return nil, ErrUserIsDisable
	}

	return claims, nil
}
