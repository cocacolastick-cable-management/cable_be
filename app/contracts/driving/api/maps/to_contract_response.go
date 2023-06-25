package maps

import (
	"fmt"
	"github.com/cable_management/cable_be/_share/errs"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/entities"
)

func ToContractRes(contract *entities.Contract) (*dtos.ContractRes, error) {

	if contract.Supplier == nil {
		return nil, errs.ErrNullReference
	}

	stock, err := contract.CalStock()
	if err != nil {
		return nil, err
	}

	isAvailable, _ := contract.IsAvailable()

	return &dtos.ContractRes{
		Id:             contract.Id,
		Counter:        contract.Counter,
		Name:           fmt.Sprintf("%v-%v", constants.ObjTyContract, contract.Counter),
		SupplierId:     contract.SupplierId,
		SupplierEmail:  contract.Supplier.Email,
		SupplierStatus: contract.Supplier.IsActive,
		CableAmount:    contract.CableAmount,
		Stock:          stock,
		StartDay:       contract.StartDay.UTC(),
		EndDay:         contract.EndDay.UTC(),
		IsAvailable:    isAvailable,
		CreatedAt:      contract.CreatedAt,
	}, nil
}
