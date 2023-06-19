package errs

import "errors"

var (
	ErrNullReference = errors.New("unexpect null reference")
)
