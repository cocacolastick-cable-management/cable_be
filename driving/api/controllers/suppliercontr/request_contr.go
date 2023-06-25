package suppliercontr

import (
	"github.com/cable_management/cable_be/app/usecases/supplier"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type RequestContr struct {
	getRequestListCase supplier.IGetSupplierRequestList
}

func NewRequestContr(getRequestListCase supplier.IGetSupplierRequestList) *RequestContr {
	return &RequestContr{getRequestListCase: getRequestListCase}
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
