package commoncontr

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestController struct {
	updateRequestStatusCase commomcase.IUpdateRequestStatus
}

func NewRequestController(updateRequestStatusCase commomcase.IUpdateRequestStatus) *RequestController {
	return &RequestController{updateRequestStatusCase: updateRequestStatusCase}
}

func (r RequestController) UpdateRequestStatus(ctx *gin.Context) {

	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)
	var reqBody = ctx.MustGet(middlewares.BodyKey).(*dtos.UpdateRequestStatusRequest)
	requestIdRaw := ctx.Param("id")

	requestId, err := uuid.Parse(requestIdRaw)
	if err != nil {
		ctx.JSON(404, types.ResponseType{
			Code:    "NF",
			Message: "request is not found",
		})
		return
	}

	requestRes, err := r.updateRequestStatusCase.Handle(accessToken, requestId, reqBody)

	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(200, types.ResponseType{
		Code:    "OK",
		Message: "update request successfully",
		Payload: requestRes,
	})
	return
}
