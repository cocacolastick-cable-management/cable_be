package plannercase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
)

type IGetContractList interface {
	Handle(accessToken string) ([]*dtos.ContractRes, error)
}

type GetContractList struct {
	authorService services.IAuthorizeService
	contractRepo  repos.IContractRepo
}

func NewGetContractList(authorService services.IAuthorizeService, contractRepo repos.IContractRepo) *GetContractList {
	return &GetContractList{authorService: authorService, contractRepo: contractRepo}
}

func (g GetContractList) Handle(accessToken string) ([]*dtos.ContractRes, error) {

	_, err := g.authorService.Authorize(accessToken, []string{constants.RolePlanner}, nil)
	if err != nil {
		return nil, err
	}

	contractList, _ := g.contractRepo.GetAll([]string{"Supplier", "RequestList"})

	contractResList := make([]*dtos.ContractRes, len(contractList))

	for i, contract := range contractList {
		contractResList[i], _ = maps.ToContractRes(contract)
	}

	return contractResList, nil
}
