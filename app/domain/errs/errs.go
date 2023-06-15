package errs

import "errors"

var (
	ErrInvalidRole               = errors.New("invalid role")
	ErrInvalidEmail              = errors.New("invalid email")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrDupEmail                  = errors.New("duplicated email")
	ErrUnauthenticated           = errors.New("unauthorized")
	ErrUnauthorized              = errors.New("authenticate failed")
	ErrUserIsDisable             = errors.New("supplier is disable")
	ErrSupplierIsDisable         = errors.New("supplier is disable")
	ErrContractorIsDisable       = errors.New("contractor is disable")
	ErrPlannerIsDisable          = errors.New("planner is disable")
	ErrInvalidJwtToken           = errors.New("invalid jwt token")
	ErrUserNotFound              = errors.New("not found user")
	ErrRequestHistoryNotFound    = errors.New("not found request history")
	ErrContractNotFound          = errors.New("not found contract")
	ErrContractorNotFound        = errors.New("not found contractor")
	ErrPlannerNotFound           = errors.New("not found planner")
	ErrUserAlreadyDisable        = errors.New("user is already disable")
	ErrUserAlreadyActive         = errors.New("user is already active")
	ErrNotIncludeRequestList     = errors.New("request list is not include")
	ErrInvalidRequestCableAmount = errors.New("invalid request cable amount")
	ErrInvalidRequestAction      = errors.New("invalid request action")
	ErrInvalidRequestStatus      = errors.New("invalid request status")
)
