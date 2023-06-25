package routers

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/suppliercontr"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
)

type SupplierRouters struct {
	requestContr  suppliercontr.IRequestContr
	contractContr suppliercontr.IContractContr
}

func NewSupplierRouters(requestContr suppliercontr.IRequestContr, contractContr suppliercontr.IContractContr) *SupplierRouters {
	return &SupplierRouters{requestContr: requestContr, contractContr: contractContr}
}

func (p SupplierRouters) Register(router gin.IRouter) {

	supplierRouter := router.Group("/supplier")

	supplierRouter.GET("/requests",
		middlewares.ParseAccessToken,
		p.requestContr.GetRequestList,
		middlewares.HandleGlobalErrors)

	supplierRouter.GET("/contracts",
		middlewares.ParseAccessToken,
		p.contractContr.GetContractList,
		middlewares.HandleGlobalErrors)
}
