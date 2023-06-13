package admincontr

import (
	"github.com/cable_management/cable_be/app/usecases/admincase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	CreateUser(ctx *gin.Context)
}

type UserController struct {
	createUserCase admincase.ICreateUser
}

func NewUserController(createUserCase admincase.ICreateUser) *UserController {
	return &UserController{createUserCase: createUserCase}
}

func (u UserController) CreateUser(ctx *gin.Context) {

	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)
	req := ctx.MustGet(middlewares.BodyKey).(*admincase.CreateUserReq)

	err := u.createUserCase.Handle(accessToken, req)
	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(201, types.ResponseType{
		Code:    "OK",
		Message: "Create user successfully",
	})
}
