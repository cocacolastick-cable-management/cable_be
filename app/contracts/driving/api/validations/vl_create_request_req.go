package validations

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/go-playground/validator/v10"
	"time"
)

type VlCreateRequestReq struct {
	contractRepo repos.IContractRepo
	userRepo     repos.IUserRepo
}

func NewVlCreateRequestReq(contractRepo repos.IContractRepo, userRepo repos.IUserRepo) *VlCreateRequestReq {
	return &VlCreateRequestReq{contractRepo: contractRepo, userRepo: userRepo}
}

func (v VlCreateRequestReq) Handle(sl validator.StructLevel) {

	req := sl.Current().Interface().(dtos.CreateRequestReq)

	// validate contractor
	contractor, _ := v.userRepo.FindByEmail(req.ContractorEmail, nil)
	if contractor == nil || contractor.Role != constants.RoleContractor {
		sl.ReportError(req.ContractorEmail, "contractorEmail", "contractorEmail", "notfound", "not found contractor")
	}
	if !contractor.IsActive {
		sl.ReportError(req.ContractorEmail, "contractorEmail", "contractorEmail", "disable-user", "contractor is disable")
	}

	// validate contract
	contract, _ := v.contractRepo.FindByCounter(req.ContractCounter, []string{"Supplier", "RequestList"})
	if contract == nil {
		sl.ReportError(req.ContractCounter, "contractCounter", "contractCounter", "notfound", "not found contract")
	}
	if contract.EndDay.Before(time.Now()) {
		sl.ReportError(req.ContractCounter, "contractCounter", "contractCounter", "expire", "contract is expired")
	}
	if !contract.Supplier.IsActive {
		sl.ReportError(req.ContractCounter, "contractCounter", "contractCounter", "unavailable", "supplier of the contract is disbale")
	}

	// validate cableAmount
	stock, _ := contract.CalStock()
	if stock < req.CableAmount {
		sl.ReportError(req.CableAmount, "cableAmount", "cableAmount", "invalid", "invalid cable amount")
	}
}
