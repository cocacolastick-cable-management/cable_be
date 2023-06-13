package middlewares

import (
	"errors"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/app/usecases/_share/errs"
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

	if errors.Is(err, services.ErrUnauthenticated) {
		ctx.JSON(401, UnauthenticatedRes)
		return
	}

	if errors.Is(err, services.ErrUnauthorized) {
		ctx.JSON(403, Unauthorized)
		return
	}

	if errors.Is(err, services.ErrUserIsDisable) {
		ctx.JSON(403, DisableAccount)
		return
	}

	if errors.Is(err, errs.ErrUserNotFound) {
		ctx.JSON(404, UserNotFound)
		return
	}

	panic(err)
}

var (
	UnauthenticatedRes = types.ResponseType{
		Code:    "UA",
		Message: "authenticate failed",
	}

	Unauthorized = types.ResponseType{
		Code:    "UA",
		Message: "unauthorized",
	}

	DisableAccount = types.ResponseType{
		Code:    "DA",
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
		Errors:  errs,
	}
}
