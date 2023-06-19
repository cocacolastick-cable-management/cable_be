package dtos

import (
	"github.com/google/uuid"
	"time"
)

type ContractRes struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	SupplierId    uuid.UUID `json:"supplierId"`
	SupplierEmail string    `json:"supplierEmail"`
	CableAmount   uint      `json:"cableAmount"`
	Stock         int       `json:"stock"`
	StartDay      time.Time `json:"startDay"`
	EndDay        time.Time `json:"endDay"`
	CreatedAt     time.Time `json:"createdAt"`
}
