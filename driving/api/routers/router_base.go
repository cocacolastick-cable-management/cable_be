package routers

import "github.com/gin-gonic/gin"

type IRouterBase interface {
	Register(router gin.IRouter)
}
