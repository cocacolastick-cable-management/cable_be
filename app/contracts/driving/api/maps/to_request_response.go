package maps

import (
	"fmt"
	"github.com/cable_management/cable_be/_share/errs"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/entities"
)

func ToRequestResponse(request *entities.Request) (*dtos.RequestRes, error) {

	if request.Contractor == nil || request.Contract == nil ||
		request.Contract.Supplier == nil || request.Planner == nil ||
		request.HistoryList == nil {
		return nil, errs.ErrNullReference
	}

	historyResList := make([]*dtos.RequestHistoryRes, len(request.HistoryList))

	for i, history := range request.HistoryList {
		historyResList[i], _ = ToRequestHistoryResponse(history)
	}

	return &dtos.RequestRes{
		Id:              request.Id,
		Name:            fmt.Sprintf("request-%v", request.Counter),
		Status:          request.Status,
		CableAmount:     request.CableAmount,
		HistoryList:     historyResList,
		ContractName:    fmt.Sprintf("contract-%v", request.Contract.Counter),
		ContractId:      request.ContractId,
		SupplierEmail:   request.Contract.Supplier.Email,
		SupplierId:      request.Contract.SupplierId,
		ContractorEmail: request.Contractor.Email,
		ContractorId:    request.ContractorId,
		PlannerEmail:    request.Planner.Email,
		PlannerId:       request.PlannerId,
		CreatedAt:       request.CreatedAt.UTC(),
	}, nil
}
