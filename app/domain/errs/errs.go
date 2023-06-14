package errs

import "errors"

var (
	ErrInvalidRole     = errors.New("invalid role")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrDupEmail        = errors.New("duplicated email")
)

var (
	ErrUnauthenticated = errors.New("authenticate failed")
	ErrUnauthorized    = errors.New("authenticate failed")
	ErrUserIsDisable   = errors.New("authenticate failed")
)

var (
	ErrInvalidJwtToken = errors.New("invalid jwt token")
)

var (
	ErrUserNotFound = errors.New("not found user")
)

var (
	ErrUserAlreadyDisable = errors.New("user is already disable")
	ErrUserAlreadyActive  = errors.New("user is already active")
)
