package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	BodyKey = "body"
)

func ParseBody[T any](ctx *gin.Context) {

	req := new(T)

	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.Set(BodyKey, req)
	ctx.Next()
}
