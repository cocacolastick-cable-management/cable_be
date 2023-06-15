package routers

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/commoncontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type CommonRouters struct {
	authContr    commoncontr.IAuthController
	requestContr commoncontr.IRequestController
}

func NewCommonRouters(authContr commoncontr.IAuthController, requestContr commoncontr.IRequestController) *CommonRouters {
	return &CommonRouters{authContr: authContr, requestContr: requestContr}
}

func (a CommonRouters) Register(router gin.IRouter) {

	commonRouter := router.Group("/common")

	commonRouter.POST("/sign-in",
		middlewares.ParseBody[dtos.SignInRequest],
		a.authContr.SignIn,
		middlewares.HandleGlobalErrors)

	commonRouter.PATCH("/requests/:id",
		middlewares.ParseAccessToken,
		middlewares.ParseBody[dtos.UpdateRequestStatusRequest],
		a.requestContr.UpdateRequestStatus,
		middlewares.HandleGlobalErrors)
}
