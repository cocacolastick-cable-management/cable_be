package commoncontr

import "github.com/gin-gonic/gin"

type IUserContr interface {
	GetUserList(ctx *gin.Context)
}
