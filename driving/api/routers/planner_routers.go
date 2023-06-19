package routers

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/plannercontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type PlannerRouters struct {
	requestContr  plannercontr.IRequestContr
	contractContr plannercontr.IContractContr
}

func NewPlannerRouters(requestContr plannercontr.IRequestContr, contractContr plannercontr.IContractContr) *PlannerRouters {
	return &PlannerRouters{requestContr: requestContr, contractContr: contractContr}
}

func (p PlannerRouters) Register(router gin.IRouter) {

	plannerRouter := router.Group("/planner")

	plannerRouter.POST("/requests",
		middlewares.ParseAccessToken,
		middlewares.ParseBody[dtos.CreateRequestReq],
		p.requestContr.CreateRequest,
		middlewares.HandleGlobalErrors)

	plannerRouter.GET("/requests",
		middlewares.ParseAccessToken,
		p.requestContr.GetRequestList,
		middlewares.HandleGlobalErrors)

	plannerRouter.GET("/contracts",
		middlewares.ParseAccessToken,
		p.contractContr.GetContractList,
		middlewares.HandleGlobalErrors)
}
