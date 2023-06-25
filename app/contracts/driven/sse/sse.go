package sse

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
	"time"
)

type NotificationDto struct {
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

type Message struct {
	ReceiverId   uuid.UUID        `json:"receiverId"`
	Notification *NotificationDto `json:"notification"`
	Object       any              `json:"object"`
}

type ISSEDriven interface {
	SendMessage(notificationList []*Message) error
}

func ToMessage(receiverId uuid.UUID, senderEmail string, notification *entities.Notification, object any) (*Message, error) {

	return &Message{
		ReceiverId: receiverId,
		Notification: &NotificationDto{
			Id:          notification.Id,
			Action:      notification.Action,
			IsRead:      notification.IsRead,
			ObjectType:  notification.ObjectType,
			ObjectId:    notification.ObjectId,
			ObjectName:  notification.ObjectName,
			SenderEmail: senderEmail,
			SenderId:    notification.SenderId,
			CreatedAt:   notification.CreatedAt,
		},
		Object: object,
	}, nil
}
