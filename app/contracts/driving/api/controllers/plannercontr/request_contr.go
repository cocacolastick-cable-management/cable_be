package plannercontr

import "github.com/gin-gonic/gin"

type IRequestContr interface {
	CreateRequest(ctx *gin.Context)
}
