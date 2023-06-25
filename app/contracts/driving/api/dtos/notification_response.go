package dtos

import (
	"github.com/google/uuid"
	"time"
)

type NotificationRes struct {
	Id uuid.UUID `json:"id"`

	Action string `json:"action"`
	IsRead bool   `json:"isRead"`

	ObjectType string    `json:"objectType"`
	ObjectId   uuid.UUID `json:"objectId"`
	ObjectName string    `json:"objectName"`

	SenderEmail string    `json:"senderEmail"`
	SenderId    uuid.UUID `json:"senderId"`

	CreatedAt time.Time `json:"createdAt"`
}
