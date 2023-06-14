package entities

import (
	"github.com/google/uuid"
	"time"
)

type Request struct {
	EntityBase

	Counter     uint   `gorm:"autoIncrement,unique"`
	Status      string `gorm:"type:varchar"`
	CableAmount uint

	ContractId   uuid.UUID
	ContractorId uuid.UUID

	Contract    *Contract         `gorm:"foreignKey:ContractId"`
	Contractor  *User             `gorm:"foreignKey:ContractorId"`
	HistoryList []*RequestHistory `gorm:"foreignKey:RequestId"`
}

func NewRequest(
	status string,
	cableAmount uint,
	contractId uuid.UUID,
	contractorId uuid.UUID,
	createdAt time.Time) *Request {

	return &Request{
		EntityBase:   NewEntityBase(createdAt),
		Status:       status,
		CableAmount:  cableAmount,
		ContractId:   contractId,
		ContractorId: contractorId}
}
