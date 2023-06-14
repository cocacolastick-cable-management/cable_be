package entities

import "github.com/google/uuid"

type RequestHistory struct {
	EntityBase

	Action string `gorm:"type:varchar"`
	Status string `gorm:"type:varchar"`

	CreatorId uuid.UUID `gorm:"type:varchar"`
	RequestId uuid.UUID `gorm:"type:varchar"`

	Creator *User    `gorm:"foreignKey:CreatorId"`
	Request *Request `gorm:"foreignKey:RequestId"`
}
