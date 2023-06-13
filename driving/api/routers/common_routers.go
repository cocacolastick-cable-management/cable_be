package routers

import (
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driving/api/controllers/commoncontr"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type CommonRouters struct {
	authContr commoncontr.IAuthController
}

func NewCommonRouters(authContr commoncontr.IAuthController) *CommonRouters {
	return &CommonRouters{authContr: authContr}
}

func (a CommonRouters) Register(router gin.IRouter) {

	router.POST("/sign-in",
		middlewares.ParseBody[commomcase.SignInRequest],
		a.authContr.SignIn,
		middlewares.HandleGlobalErrors)
}
