package errs

import "errors"

var (
	ErrInvalidRole           = errors.New("invalid role")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrDupEmail              = errors.New("duplicated email")
	ErrUnauthenticated       = errors.New("authenticate failed")
	ErrUnauthorized          = errors.New("authenticate failed")
	ErrUserIsDisable         = errors.New("authenticate failed")
	ErrInvalidJwtToken       = errors.New("invalid jwt token")
	ErrUserNotFound          = errors.New("not found user")
	ErrUserAlreadyDisable    = errors.New("user is already disable")
	ErrUserAlreadyActive     = errors.New("user is already active")
	ErrNotIncludeRequestList = errors.New("request list is not include")
)
