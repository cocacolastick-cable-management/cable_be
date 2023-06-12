package commom_case

import (
	"errors"
	"github.com/cable_management/cable_be/app/contracts/database/repos"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/share/errs"
)

var (
	ErrUnauthenticated = errors.New("authenticate failed")
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Email     string              `json:"email"`
	Role      string              `json:"role"`
	Name      string              `json:"name"`
	AuthToken *services.AuthToken `json:"authToken"`
}

type ISignIn interface {
	Handle(req *SignInRequest) (res *SignInResponse, err error)
}

type SignIn struct {
	userRepo        repos.IUserRepo
	tokenService    services.IAuthTokenService
	passwordService services.IPasswordService
}

func (s SignIn) Handle(req *SignInRequest) (res *SignInResponse, err error) {

	user, _ := s.userRepo.FindByEmail(req.Email, nil)
	if user == nil || !s.passwordService.Compare(req.Password, user.PasswordHash) {
		return nil, ErrUnauthenticated
	}

	authToken, err := s.tokenService.CreateAuthToken(user, nil)
	if err != nil {
		return nil, err
	}

	return toSignInRes(user, authToken)
}

func toSignInRes(user *entities.User, authToken *services.AuthToken) (*SignInResponse, error) {

	if user == nil || authToken == nil {
		return nil, errs.ErrNullException
	}

	res := &SignInResponse{
		Email:     user.Email,
		Role:      user.Role,
		Name:      user.Name,
		AuthToken: authToken,
	}

	return res, nil
}
