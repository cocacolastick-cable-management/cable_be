package errs

import "errors"

var (
	ErrNullException = errors.New("unexpect null variable")
)
