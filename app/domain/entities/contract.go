package entities

import (
	sherrs "github.com/cable_management/cable_be/_share/errs"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/google/uuid"
	"time"
)

type Contract struct {
	EntityBase

	Counter uint `gorm:"autoIncrement,unique"`

	CableAmount uint
	StartDay    time.Time
	EndDay      time.Time

	SupplierId uuid.UUID

	Supplier    *User      `gorm:"foreignKey:SupplierId"`
	RequestList []*Request `gorm:"foreignKey:ContractId"`
}

func (c Contract) CalStock() (stock uint, err error) {

	if c.RequestList == nil {
		return 0, errs.ErrNotIncludeRequestList
	}

	stock = c.CableAmount

	for _, request := range c.RequestList {
		stock -= request.CableAmount
	}

	return stock, nil
}

func (c Contract) IsAvailable() (bool, error) {

	if c.Supplier == nil {
		return false, sherrs.ErrNullReference
	}

	return c.Supplier.IsActive && c.EndDay.Before(time.Now()), nil
}
