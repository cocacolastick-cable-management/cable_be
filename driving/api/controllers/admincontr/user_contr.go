package admincontr

import (
	"errors"
	"github.com/cable_management/cable_be/app/usecases/_share/errs"
	"github.com/cable_management/cable_be/app/usecases/admincase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserController interface {
	CreateUser(ctx *gin.Context)
	UpdateUserIsActive(ctx *gin.Context)
}

type UserController struct {
	createUserCase         admincase.ICreateUser
	updateUserIsActiveCase admincase.IUpdateUserIsActive
}

func NewUserController(createUserCase admincase.ICreateUser, updateUserIsActive admincase.IUpdateUserIsActive) *UserController {
	return &UserController{createUserCase: createUserCase, updateUserIsActiveCase: updateUserIsActive}
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

func (u UserController) UpdateUserIsActive(ctx *gin.Context) {

	// parse request
	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)
	req := ctx.MustGet(middlewares.BodyKey).(*admincase.UpdateUserIsActiveReq)
	userIdRaw := ctx.Param("id")

	userId, err := uuid.Parse(userIdRaw)
	if err != nil {
		ctx.Set(constants.ErrKey, errs.ErrUserNotFound)
		ctx.Next()
		return
	}

	// execute
	err = u.updateUserIsActiveCase.Handle(accessToken, userId, req)

	// handle error
	if err != nil {

		switch {
		case errors.Is(err, admincase.ErrUserAlreadyActive):
			ctx.JSON(400, types.ResponseType{
				Code:    "AA",
				Message: "user is already active",
			})
			return
		case errors.Is(err, admincase.ErrUserAlreadyDisable):
			ctx.JSON(400, types.ResponseType{
				Code:    "AD",
				Message: "user is already disable",
			})
			return
		}

		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(200, types.ResponseType{
		Code:    "OK",
		Message: "update user successfully",
	})
	return
}
