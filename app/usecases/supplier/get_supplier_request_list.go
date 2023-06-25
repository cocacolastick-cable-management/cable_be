package supplier

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
)

type IGetSupplierRequestList interface {
	Handle(accessToken string) ([]*dtos.RequestRes, error)
}

type GetSupplierRequestList struct {
	requestRepo   repos.IRequestRepo
	authorService services.IAuthorizeService
}

func NewGetSupplierRequestList(requestRepo repos.IRequestRepo, authorService services.IAuthorizeService) *GetSupplierRequestList {
	return &GetSupplierRequestList{requestRepo: requestRepo, authorService: authorService}
}

func (g GetSupplierRequestList) Handle(accessToken string) ([]*dtos.RequestRes, error) {

	claims, err := g.authorService.Authorize(accessToken, []string{constants.RoleSupplier}, nil)
	if err != nil {
		return nil, err
	}

	requestList, _ := g.requestRepo.GetBySupplierId(claims.UserId,
		[]string{"Contractor", "Contract", "Contract.Supplier", "Planner", "HistoryList", "HistoryList.Creator"})

	requestResList := make([]*dtos.RequestRes, len(requestList))
	for i, request := range requestList {
		requestResList[i], _ = maps.ToRequestResponse(request)
	}

	return requestResList, nil
}
