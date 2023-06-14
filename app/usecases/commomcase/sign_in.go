package commomcase

import (
	shErrs "github.com/cable_management/cable_be/_share/errs"
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/domain/services"
)

type ISignIn interface {
	Handle(req *dtos.SignInRequest) (res *dtos.SignInRes, err error)
}

type SignIn struct {
	userRepo        repos.IUserRepo
	tokenService    services.IAuthTokenService
	passwordService services.IPasswordService
}

func NewSignIn(
	userRepo repos.IUserRepo,
	tokenService services.IAuthTokenService,
	passwordService services.IPasswordService) *SignIn {

	return &SignIn{
		userRepo:        userRepo,
		tokenService:    tokenService,
		passwordService: passwordService}
}

func (s SignIn) Handle(req *dtos.SignInRequest) (res *dtos.SignInRes, err error) {

	user, _ := s.userRepo.FindByEmail(req.Email, nil)
	if user == nil {
		return nil, errs.ErrUnauthenticated
	}

	if !s.passwordService.Compare(req.Password, user.PasswordHash) {
		return nil, errs.ErrUnauthenticated
	}

	authToken, err := s.tokenService.CreateAuthToken(user, nil)
	if err != nil {
		return nil, err
	}

	return toSignInRes(user, authToken)
}

func toSignInRes(user *entities.User, authToken *services.AuthToken) (*dtos.SignInRes, error) {

	if user == nil || authToken == nil {
		return nil, shErrs.ErrNullException
	}

	res := &dtos.SignInRes{
		Email:     user.Email,
		Role:      user.Role,
		Name:      user.Name,
		AuthToken: authToken,
	}

	return res, nil
}
