package dtos

import (
	"github.com/google/uuid"
	"time"
)

type RequestRes struct {
	Id uuid.UUID `json:"id"`

	Name        string               `json:"name"`
	Status      string               `json:"status"`
	CableAmount uint                 `json:"cableAmount"`
	HistoryList []*RequestHistoryRes `json:"historyList"`

	ContractName string    `json:"contractName"`
	ContractId   uuid.UUID `json:"contractId"`

	SupplierEmail string    `json:"supplierEmail"`
	SupplierId    uuid.UUID `json:"supplierId"`

	ContractorEmail string    `json:"contractorEmail"`
	ContractorId    uuid.UUID `json:"contractorId"`

	PlannerEmail string    `json:"plannerEmail"`
	PlannerId    uuid.UUID `json:"plannerId"`

	CreatedAt time.Time `json:"createdAt"`
}
