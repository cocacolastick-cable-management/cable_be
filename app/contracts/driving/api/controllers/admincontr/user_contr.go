package admincontr

import "github.com/gin-gonic/gin"

type IUserController interface {
	CreateUser(ctx *gin.Context)
	UpdateUserIsActive(ctx *gin.Context)
}
