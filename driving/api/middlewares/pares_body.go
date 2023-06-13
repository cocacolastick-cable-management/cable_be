package middlewares

import (
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	BodyKey = "body"
)

func ParseBody[T any](ctx *gin.Context) {

	req := new(T)

	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ResponseType{
			Code:    "BR",
			Message: "bad request",
			Errors:  err,
		})
	}

	ctx.Set(BodyKey, req)
	ctx.Next()
}
