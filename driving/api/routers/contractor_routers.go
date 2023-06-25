package routers

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/contractorcontr"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type ContractorRouters struct {
	requestContr contractorcontr.IRequestContr
}

func NewContractorRouters(requestContr contractorcontr.IRequestContr) *ContractorRouters {
	return &ContractorRouters{requestContr: requestContr}
}

func (c ContractorRouters) Register(router gin.IRouter) {

	contractorRouter := router.Group("/contractor")

	contractorRouter.GET("/requests",
		middlewares.ParseAccessToken,
		c.requestContr.GetRequestList,
		middlewares.HandleGlobalErrors)
}
