package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
)

type INotificationRepo interface {
	InsertMany(notificationList []*entities.Notification) error
	FindByReceiverId(receiverId uuid.UUID, withs []string) ([]*entities.Notification, error)
	Save(notification *entities.Notification) error
	FindById(id uuid.UUID, withs []string) (*entities.Notification, error)
}
