package commoncontr

import "github.com/gin-gonic/gin"

type IAuthController interface {
	SignIn(ctx *gin.Context)
}
