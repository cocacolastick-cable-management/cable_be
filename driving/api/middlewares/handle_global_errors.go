package middlewares

import (
	"errors"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleGlobalErrors(ctx *gin.Context) {

	errRaw, ok := ctx.Get(constants.ErrKey)
	if !ok {
		// TODO bad request
		return
	}

	err, ok := errRaw.(error)
	if !ok {
		// TODO bad request
		return
	}

	if validErrors, ok := err.(validator.ValidationErrors); ok {
		ctx.JSON(400, toValidationErrorRes(validErrors))
		return
	}

	if errors.Is(err, errs.ErrUnauthenticated) {
		ctx.JSON(401, Unauthenticated)
		return
	}

	if errors.Is(err, errs.ErrUnauthorized) {
		ctx.JSON(403, Unauthorized)
		return
	}

	if errors.Is(err, errs.ErrUserIsDisable) {
		ctx.JSON(403, DisableAccount)
		return
	}

	if errors.Is(err, errs.ErrUserNotFound) {
		ctx.JSON(404, UserNotFound)
		return
	}

	if errors.Is(err, errs.ErrInvalidRequestStatus) {
		ctx.JSON(400, types.ResponseType{
			Code:    "BR",
			Message: "invalid request",
		})
		return
	}

	if errors.Is(err, errs.ErrRequestNotFound) {
		ctx.JSON(404, types.ResponseType{
			Code:    "NF",
			Message: "request is not found",
		})
		return
	}

	if errors.Is(err, errs.ErrContractNotFound) {
		ctx.JSON(400, types.ResponseType{
			Code:    "CU",
			Message: "contract is unavailable",
		})
		return
	}

	if errors.Is(err, errs.ErrNotificationNotFound) {
		ctx.JSON(400, types.ResponseType{
			Code:    "NNF",
			Message: "notification is not found",
		})
		return
	}

	panic(err)
}

var (
	Unauthenticated = types.ResponseType{
		Code:    constants.CodeAF,
		Message: "authenticate failed",
	}

	Unauthorized = types.ResponseType{
		Code:    "UA",
		Message: "unauthorized",
	}

	DisableAccount = types.ResponseType{
		Code:    constants.CodeDA,
		Message: "account is disable",
	}

	UserNotFound = types.ResponseType{
		Code:    "NF",
		Message: "user is not found",
	}
)

func toValidationErrorRes(errs validator.ValidationErrors) types.ResponseType {
	// TODO: map to a good format
	return types.ResponseType{
		Code:    "IV",
		Message: "invalid fields",
		Errors:  ValidationErrorToStruct(errs),
	}
}

type ValidationErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidationErrorToStruct(err validator.ValidationErrors) []*ValidationErrorResponse {
	var errRes []*ValidationErrorResponse
	if err != nil {
		for _, err := range err {
			var element ValidationErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errRes = append(errRes, &element)
		}
	}
	return errRes
}
