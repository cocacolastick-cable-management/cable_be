package contractor

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
)

type IGetContractorRequestList interface {
	Handle(accessToken string) ([]*dtos.RequestRes, error)
}

type GetContractorRequestList struct {
	requestRepo   repos.IRequestRepo
	authorService services.IAuthorizeService
}

func NewGetContractorRequestList(requestRepo repos.IRequestRepo, authorService services.IAuthorizeService) *GetContractorRequestList {
	return &GetContractorRequestList{requestRepo: requestRepo, authorService: authorService}
}

func (g GetContractorRequestList) Handle(accessToken string) ([]*dtos.RequestRes, error) {

	claims, err := g.authorService.Authorize(accessToken, []string{constants.RoleContractor}, nil)
	if err != nil {
		return nil, err
	}

	requestList, _ := g.requestRepo.GetByContractorId(claims.UserId,
		[]string{"Contractor", "Contract", "Contract.Supplier", "Planner", "HistoryList", "HistoryList.Creator"})

	requestResList := make([]*dtos.RequestRes, len(requestList))
	for i, request := range requestList {
		requestResList[i], _ = maps.ToRequestResponse(request)
	}

	return requestResList, nil
}
