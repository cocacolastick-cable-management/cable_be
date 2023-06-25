package commoncontr

import (
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserContr struct {
	getUserListCase commomcase.IGetUserList
}

func NewUserContr(getUserListCase commomcase.IGetUserList) *UserContr {
	return &UserContr{getUserListCase: getUserListCase}
}

func (u UserContr) GetUserList(ctx *gin.Context) {

	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)
	roles := strings.Split(ctx.Query("roles"), ",")

	userResList, err := u.getUserListCase.Handle(accessToken, roles)

	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(200, types.ResponseType{
		Code:    "OK",
		Message: "OK",
		Payload: userResList,
	})
	return
}
