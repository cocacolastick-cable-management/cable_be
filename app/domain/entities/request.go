package entities

import (
	"github.com/google/uuid"
)

type Request struct {
	EntityBase

	Counter     uint   `gorm:"autoIncrement"`
	Status      string `gorm:"type:varchar"`
	CableAmount uint

	ContractId   uuid.UUID
	ContractorId uuid.UUID

	Contract   *Contract `gorm:"foreignKey:ContractId"`
	Contractor *User     `gorm:"foreignKey:ContractorId"`
}
