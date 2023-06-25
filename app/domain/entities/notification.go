package entities

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	EntityBase

	Action     string `gorm:"type:varchar"`
	IsRead     bool
	ObjectType string    `gorm:"type:varchar"`
	ObjectId   uuid.UUID `gorm:"type:varchar"`
	ObjectName string    `gorm:"type:varchar"`

	SenderId   uuid.UUID
	ReceiverId uuid.UUID

	Sender   *User `gorm:"foreignKey:SenderId"`
	Receiver *User `gorm:"foreignKey:ReceiverId"`

	Object any `gorm:"-:all"`
	//Object any `gorm:"<-:false"`
}

func NewNotification(action string, isRead bool, objectType string, objectId uuid.UUID, objectName string, createdAt time.Time, senderId uuid.UUID, receiverId uuid.UUID) *Notification {
	return &Notification{EntityBase: NewEntityBase(createdAt), Action: action, IsRead: isRead, ObjectType: objectType, ObjectId: objectId, ObjectName: objectName, SenderId: senderId, ReceiverId: receiverId}
}
