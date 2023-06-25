package services

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const (
	AccessTokenExpire  = time.Hour * 2
	RefreshTokenExpire = time.Hour * 24 * 30
)

const (
	PermResetPassword = "update:password"
)

const (
	AccessTokenTypeName  = "access-token"
	RefreshTokenTypeName = "refresh-token"
)

type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthTokenClaims struct {
	jwt.RegisteredClaims
	Role        string    `json:"role"`
	UserEmail   string    `json:"email"`
	Type        string    `json:"type"`
	UserId      uuid.UUID `json:"user_id"`
	Permissions []string  `json:"permissions,omitempty"`
}

type IAuthTokenService interface {
	CreateAuthToken(user *entities.User, permissions []string) (authToken *AuthToken, err error)
	IsAccessTokenValid(accessToken string) (bool, *AuthTokenClaims)
	IsRefreshTokenValid(refreshToken string) (bool, *AuthTokenClaims)
}

type AuthTokenService struct {
	jwtSecret string
}

func NewAuthTokenService(jwtSecret string) *AuthTokenService {
	return &AuthTokenService{jwtSecret: jwtSecret}
}

func (a AuthTokenService) CreateAuthToken(user *entities.User, permissions []string) (authToken *AuthToken, err error) {

	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		AuthTokenClaims{
			Role:        user.Role,
			UserEmail:   user.Email,
			UserId:      user.Id,
			Type:        AccessTokenTypeName,
			Permissions: permissions,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
			},
		})

	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		AuthTokenClaims{
			Role:      user.Role,
			UserEmail: user.Email,
			UserId:    user.Id,
			Type:      RefreshTokenTypeName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpire)),
			},
		})

	accessTokenStr, _ := accessToken.SignedString([]byte(a.jwtSecret))
	refreshTokenStr, _ := refreshToken.SignedString([]byte(a.jwtSecret))

	return &AuthToken{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}, nil
}

func (a AuthTokenService) IsAccessTokenValid(accessToken string) (bool, *AuthTokenClaims) {

	claims, err := a.ParseToClaims(accessToken)
	if err != nil {
		return false, nil
	}

	tokenType := claims.Type
	if tokenType != AccessTokenTypeName {
		return false, nil
	}

	return true, claims
}

func (a AuthTokenService) IsRefreshTokenValid(refreshToken string) (bool, *AuthTokenClaims) {

	claims, err := a.ParseToClaims(refreshToken)
	if err != nil {
		return false, nil
	}

	tokenType := claims.Type
	if tokenType != RefreshTokenTypeName {
		return false, nil
	}

	return true, claims
}

func (a AuthTokenService) ParseToClaims(tokenStr string) (*AuthTokenClaims, error) {

	claims := &AuthTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errs.ErrInvalidJwtToken
	}

	return claims, nil
}
