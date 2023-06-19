package routers

import (
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type IRouterBase interface {
	Register(router gin.IRouter)
}

type RouterBase struct {
}

func NewRouterBase() *RouterBase {
	return &RouterBase{}
}

func (r RouterBase) Register(router gin.IRouter) {
	router.Use(middlewares.Cors)
}
