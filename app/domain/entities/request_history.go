package entities

import (
	"github.com/google/uuid"
	"time"
)

type RequestHistory struct {
	EntityBase

	Action string `gorm:"type:varchar"`
	Status string `gorm:"type:varchar"`

	CreatorId uuid.UUID `gorm:"type:varchar"`
	RequestId uuid.UUID `gorm:"type:varchar"`

	Creator *User    `gorm:"foreignKey:CreatorId"`
	Request *Request `gorm:"foreignKey:RequestId"`
}

func NewRequestHistory(
	action string,
	status string,
	creatorId uuid.UUID,
	requestId uuid.UUID,
	createdAt time.Time) *RequestHistory {

	return &RequestHistory{
		EntityBase: NewEntityBase(createdAt),
		Action:     action,
		Status:     status,
		CreatorId:  creatorId,
		RequestId:  requestId}
}
