package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	db *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) *NotificationRepo {
	return &NotificationRepo{db: db}
}

func (n NotificationRepo) InsertMany(notificationList []*entities.Notification) error {
	result := n.db.Create(notificationList)
	return result.Error
}

func (n NotificationRepo) Save(notification *entities.Notification) error {
	result := n.db.Save(notification)
	return result.Error
}

func (n NotificationRepo) FindByReceiverId(receiverId uuid.UUID, withs []string) ([]*entities.Notification, error) {

	var notificationList []*entities.Notification
	query := n.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.
		Order("notifications.created_at desc").
		Find(&notificationList, "notifications.receiver_id = ?", receiverId)
	if result.Error != nil {
		return nil, result.Error
	}

	return notificationList, nil
}

func (n NotificationRepo) FindById(id uuid.UUID, withs []string) (*entities.Notification, error) {

	query := n.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	var notification *entities.Notification
	result := query.First(&notification, "notifications.id = ?", id)

	return notification, result.Error
}
