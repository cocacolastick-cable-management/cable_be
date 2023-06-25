package commoncontr

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotificationContr struct {
	getNotificationListCase      commomcase.IGetNotificationList
	updateNotificationIsReadCase commomcase.IUpdateNotificationIsRead
}

func NewNotificationContr(getNotificationListCase commomcase.IGetNotificationList, updateNotificationIsReadCase commomcase.IUpdateNotificationIsRead) *NotificationContr {
	return &NotificationContr{getNotificationListCase: getNotificationListCase, updateNotificationIsReadCase: updateNotificationIsReadCase}
}

func (n NotificationContr) GetNotificationList(ctx *gin.Context) {

	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)

	notificationResList, err := n.getNotificationListCase.Handle(accessToken)

	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(200, types.ResponseType{
		Code:    "OK",
		Message: "OK",
		Payload: notificationResList,
	})
	return
}

func (n NotificationContr) UpdateNotificationIsRead(ctx *gin.Context) {

	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)
	req := ctx.MustGet(middlewares.BodyKey).(*dtos.UpdateNotificationIsReadReq)
	notiIdRaw := ctx.Param("id")

	notiId, err := uuid.Parse(notiIdRaw)
	if err != nil {
		ctx.Set(constants.ErrKey, errs.ErrNotificationNotFound)
		ctx.Next()
		return
	}

	err = n.updateNotificationIsReadCase.Handle(accessToken, notiId, req)

	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(200, types.ResponseType{
		Code:    "OK",
		Message: "OK",
	})
	return
}
