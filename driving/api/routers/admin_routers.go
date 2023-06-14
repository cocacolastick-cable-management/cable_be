package routers

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/admincontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type AdminRouters struct {
	userContr admincontr.IUserController
}

func NewAdminRouters(userContr admincontr.IUserController) *AdminRouters {
	return &AdminRouters{userContr: userContr}
}

func (a AdminRouters) Register(router gin.IRouter) {

	adminRouter := router.Group("/admin")

	adminRouter.POST("/users",
		middlewares.ParseAccessToken,
		middlewares.ParseBody[dtos.CreateUserReq],
		a.userContr.CreateUser,
		middlewares.HandleGlobalErrors)

	adminRouter.PATCH("/users/:id",
		middlewares.ParseAccessToken,
		middlewares.ParseBody[dtos.UpdateUserIsActiveReq],
		a.userContr.UpdateUserIsActive,
		middlewares.HandleGlobalErrors)
}
