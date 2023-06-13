package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	AccessTokenKey = "accessToken"
)

func ParseAccessToken(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		ctx.Set(AccessTokenKey, accessToken)
		ctx.Next()
	} else {
		ctx.JSON(401, UnauthenticatedRes)
	}
}
