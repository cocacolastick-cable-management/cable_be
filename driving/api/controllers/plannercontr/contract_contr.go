package plannercontr

import (
	"github.com/cable_management/cable_be/app/usecases/plannercase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/types"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type ContractContr struct {
	getContractListCase plannercase.IGetContractList
}

func NewContractContr(getContractListCase plannercase.IGetContractList) *ContractContr {
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
