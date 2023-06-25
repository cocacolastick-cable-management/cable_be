package commoncontr

import "github.com/gin-gonic/gin"

type INotificationContr interface {
	GetNotificationList(ctx *gin.Context)
	UpdateNotificationIsRead(ctx *gin.Context)
}
