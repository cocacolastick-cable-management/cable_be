package routers

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/commoncontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type CommonRouters struct {
	authContr         commoncontr.IAuthController
	requestContr      commoncontr.IRequestController
	userContr         commoncontr.IUserContr
	notificationContr commoncontr.INotificationContr
}

func NewCommonRouters(authContr commoncontr.IAuthController, requestContr commoncontr.IRequestController, userContr commoncontr.IUserContr, notificationContr commoncontr.INotificationContr) *CommonRouters {
	return &CommonRouters{authContr: authContr, requestContr: requestContr, userContr: userContr, notificationContr: notificationContr}
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

	commonRouter.GET("/users",
		middlewares.ParseAccessToken,
		a.userContr.GetUserList,
		middlewares.HandleGlobalErrors)

	commonRouter.GET("/notifications",
		middlewares.ParseAccessToken,
		a.notificationContr.GetNotificationList,
		middlewares.HandleGlobalErrors)

	commonRouter.PATCH("/notifications/:id",
		middlewares.ParseAccessToken,
		middlewares.ParseBody[dtos.UpdateNotificationIsReadReq],
		a.notificationContr.UpdateNotificationIsRead,
		middlewares.HandleGlobalErrors)
}
