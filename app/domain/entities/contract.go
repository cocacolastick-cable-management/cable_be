package entities

import (
	"github.com/google/uuid"
	"time"
)

type Contract struct {
	EntityBase

	Counter uint `gorm:"autoIncrement"`

	CableAmount uint
	StartDay    time.Time
	EndDay      time.Time

	SupplierId uuid.UUID

	Supplier *User      `gorm:"foreignKey:SupplierId"`
	Requests []*Request `gorm:"foreignKey:ContractId"`
}
