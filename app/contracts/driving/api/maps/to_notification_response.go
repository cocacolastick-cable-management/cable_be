package maps

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/entities"
)

func ToNotificationResponse(notification *entities.Notification) (*dtos.NotificationRes, error) {

	return &dtos.NotificationRes{
		Id:          notification.Id,
		Action:      notification.Action,
		IsRead:      notification.IsRead,
		ObjectType:  notification.ObjectType,
		ObjectId:    notification.ObjectId,
		ObjectName:  notification.ObjectName,
		SenderEmail: notification.Sender.Email,
		SenderId:    notification.SenderId,
		CreatedAt:   notification.CreatedAt,
	}, nil
}
