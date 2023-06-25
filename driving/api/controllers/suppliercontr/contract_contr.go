package suppliercontr

import (
	"github.com/cable_management/cable_be/app/usecases/supplier"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type ContractContr struct {
	getContractListCase supplier.IGetSupplierContractList
}

func NewContractContr(getContractListCase supplier.IGetSupplierContractList) *ContractContr {
	return &ContractContr{getContractListCase: getContractListCase}
}

func (c ContractContr) GetContractList(ctx *gin.Context) {

	accessToken := ctx.MustGet(middlewares.AccessTokenKey).(string)

	contractList, err := c.getContractListCase.Handle(accessToken)

	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(200, types.ResponseType{
		Code:    "OK",
		Message: "OK",
		Payload: contractList,
	})
	return
}
