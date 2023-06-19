package plannercontr

import "github.com/gin-gonic/gin"

type IContractContr interface {
	GetContractList(ctx *gin.Context)
}
