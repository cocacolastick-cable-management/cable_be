package commomcase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/services"
)

type IGetNotificationList interface {
	Handle(accessToken string) ([]*dtos.NotificationRes, error)
}

type GetNotificationList struct {
	notificationRepo repos.INotificationRepo
	authorService    services.IAuthorizeService
}

func NewGetNotificationList(notificationRepo repos.INotificationRepo, authorService services.IAuthorizeService) *GetNotificationList {
	return &GetNotificationList{notificationRepo: notificationRepo, authorService: authorService}
}

func (g GetNotificationList) Handle(accessToken string) ([]*dtos.NotificationRes, error) {

	claims, err := g.authorService.Authorize(accessToken, nil, nil)
	if err != nil {
		return nil, err
	}

	notificationList, _ := g.notificationRepo.FindByReceiverId(claims.UserId, []string{"Sender"})

	notificationResList := make([]*dtos.NotificationRes, len(notificationList))
	for i, notification := range notificationList {
		notificationResList[i], _ = maps.ToNotificationResponse(notification)
	}

	return notificationResList, nil
}
