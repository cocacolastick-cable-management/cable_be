package supplier

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
)

type IGetSupplierContractList interface {
	Handle(accessToken string) ([]*dtos.ContractRes, error)
}

type GetSupplierContractList struct {
	contractRepo  repos.IContractRepo
	authorService services.IAuthorizeService
}

func NewGetSupplierContractList(contractRepo repos.IContractRepo, authorService services.IAuthorizeService) *GetSupplierContractList {
	return &GetSupplierContractList{contractRepo: contractRepo, authorService: authorService}
}

func (g GetSupplierContractList) Handle(accessToken string) ([]*dtos.ContractRes, error) {

	claims, err := g.authorService.Authorize(accessToken, []string{constants.RoleSupplier}, nil)
	if err != nil {
		return nil, err
	}

	contractList, _ := g.contractRepo.GetBySupplierId(claims.UserId, []string{"Supplier", "RequestList"})

	contractResList := make([]*dtos.ContractRes, len(contractList))

	for i, contract := range contractList {
		contractResList[i], _ = maps.ToContractRes(contract)
	}

	return contractResList, nil
}
