package middlewares

import (
	"errors"
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/responses"
	"github.com/gin-gonic/gin"
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

	if errors.Is(err, commomcase.ErrUnauthenticated) {
		ctx.JSON(401, responses.UnauthenticatedRes)
		return
	}

	panic(err)
}
