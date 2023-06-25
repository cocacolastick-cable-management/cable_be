package suppliercontr

import "github.com/gin-gonic/gin"

type IRequestContr interface {
	GetRequestList(ctx *gin.Context)
}
