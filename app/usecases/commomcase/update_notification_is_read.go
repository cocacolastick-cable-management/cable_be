package commomcase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/google/uuid"
)

type IUpdateNotificationIsRead interface {
	Handle(accessToken string, notificationId uuid.UUID, req *dtos.UpdateNotificationIsReadReq) error
}

type UpdateNotificationIsRead struct {
	authorService    services.IAuthorizeService
	notificationRepo repos.INotificationRepo
}

func NewUpdateNotificationIsRead(authorService services.IAuthorizeService, notificationRepo repos.INotificationRepo) *UpdateNotificationIsRead {
	return &UpdateNotificationIsRead{authorService: authorService, notificationRepo: notificationRepo}
}

func (u UpdateNotificationIsRead) Handle(accessToken string, notificationId uuid.UUID, req *dtos.UpdateNotificationIsReadReq) error {

	claims, err := u.authorService.Authorize(accessToken, nil, nil)
	if err != nil {
		return err
	}

	notification, _ := u.notificationRepo.FindById(notificationId, nil)
	if claims.UserId != notification.ReceiverId {
		return errs.ErrNotificationNotFound
	}

	notification.IsRead = *req.IsRead
	_ = u.notificationRepo.Save(notification)
	return nil
}
