package plannercontr

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/usecases/plannercase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type RequestContr struct {
	createRequestCase  plannercase.ICreateRequest
	getRequestListCase plannercase.IGetRequestList
}

func NewRequestContr(createRequestCase plannercase.ICreateRequest, getRequestListCase plannercase.IGetRequestList) *RequestContr {
	return &RequestContr{createRequestCase: createRequestCase, getRequestListCase: getRequestListCase}
}

func (r RequestContr) CreateRequest(ctx *gin.Context) {

	// parse request
	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)
	req := ctx.MustGet(middlewares.BodyKey).(*dtos.CreateRequestReq)

	// excute
	err := r.createRequestCase.Handle(accessToken, req)

	// handle error
	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(201, types.ResponseType{
		Code:    "OK",
		Message: "Create request successfully",
	})
	return
}

func (r RequestContr) GetRequestList(ctx *gin.Context) {

	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)

	requestList, err := r.getRequestListCase.Handle(accessToken)

	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(200, types.ResponseType{
		Code:    "OK",
		Message: "OK",
		Payload: requestList,
	})
	return
}
