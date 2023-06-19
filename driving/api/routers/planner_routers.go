package routers

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/plannercontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type PlannerRouters struct {
	requestContr plannercontr.IRequestContr
}

func NewPlannerRouters(requestContr plannercontr.IRequestContr) *PlannerRouters {
	return &PlannerRouters{requestContr: requestContr}
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
}
