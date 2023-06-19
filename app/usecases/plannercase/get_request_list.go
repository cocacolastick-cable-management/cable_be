package plannercase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
)

type IGetRequestList interface {
	Handle(accessToken string) ([]*dtos.RequestRes, error)
}

type GetRequestList struct {
	requestRepo   repos.IRequestRepo
	authorService services.IAuthorizeService
}

func NewGetRequestList(requestRepo repos.IRequestRepo, authorService services.IAuthorizeService) *GetRequestList {
	return &GetRequestList{requestRepo: requestRepo, authorService: authorService}
}

func (g GetRequestList) Handle(accessToken string) ([]*dtos.RequestRes, error) {

	_, err := g.authorService.Authorize(accessToken, []string{constants.RolePlanner}, nil)
	if err != nil {
		return nil, err
	}

	requestList, _ := g.requestRepo.GetAll([]string{"Contractor", "Contract", "Contract.Supplier", "Planner", "HistoryList", "HistoryList.Creator"})

	requestResList := make([]*dtos.RequestRes, len(requestList))
	for i, request := range requestList {
		requestResList[i], _ = maps.ToRequestResponse(request)
	}

	return requestResList, nil
}
