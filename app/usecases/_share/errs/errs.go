package errs

import "errors"

var (
	ErrUserNotFound = errors.New("not found user")
)
