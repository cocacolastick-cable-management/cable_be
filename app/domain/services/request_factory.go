package services

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/errs"
	"time"
)

type IRequestFactory interface {
	CreateRequest(cableAmount uint, contractCounter uint, contractorEmail string) (request *entities.Request, err error)
}

type RequestFactory struct {
	contractRepo repos.IContractRepo
	userRepo     repos.IUserRepo
}

func NewRequestFactory(contractRepo repos.IContractRepo, userRepo repos.IUserRepo) *RequestFactory {
	return &RequestFactory{contractRepo: contractRepo, userRepo: userRepo}
}

func (r RequestFactory) CreateRequest(cableAmount uint, contractCounter uint, contractorEmail string) (request *entities.Request, err error) {

	// find contractor
	contractor, _ := r.userRepo.FindByEmail(contractorEmail, nil)
	if contractor == nil || contractor.Role != constants.RoleContractor {
		return nil, errs.ErrContractorNotFound
	}
	if !contractor.IsActive {
		return nil, errs.ErrContractorIsDisable
	}

	// find contract
	contract, _ := r.contractRepo.FindByCounter(contractCounter, []string{"Supplier", "RequestList"})
	if contract == nil {
		return nil, errs.ErrContractNotFound
	}
	if !contract.Supplier.IsActive {
		return nil, errs.ErrSupplierIsDisable
	}

	// validate cableAmount
	stock, _ := contract.CalStock()
	if stock < cableAmount {
		return nil, errs.ErrInvalidRequestCableAmount
	}

	// create request
	return entities.NewRequest(
		constants.StatusNew,
		cableAmount,
		contract.Id,
		contractor.Id,
		time.Now()), nil
}
