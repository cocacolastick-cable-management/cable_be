package commoncontr

import "github.com/gin-gonic/gin"

type IRequestController interface {
	UpdateRequestStatus(ctx *gin.Context)
}
