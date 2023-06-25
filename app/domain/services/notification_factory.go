package services

import (
	"fmt"
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/google/uuid"
	"time"
)

type INotificationFactory interface {
	CreateNotificationListForRequestAction(historyId uuid.UUID) ([]*entities.Notification, error)
	CreateNotificationListForUserAction(senderId uuid.UUID, user *entities.User, action string) ([]*entities.Notification, error)
}

type NotificationFactory struct {
	userRepo           repos.IUserRepo
	requestHistoryRepo repos.IRequestHistoryRepo
}

func NewNotificationFactory(userRepo repos.IUserRepo, requestHistoryRepo repos.IRequestHistoryRepo) *NotificationFactory {
	return &NotificationFactory{userRepo: userRepo, requestHistoryRepo: requestHistoryRepo}
}

func (n NotificationFactory) CreateNotificationListForRequestAction(historyId uuid.UUID) ([]*entities.Notification, error) {

	history, _ := n.requestHistoryRepo.FindById(historyId,
		[]string{"Creator", "Request", "Request.Planner", "Request.Contractor", "Request.Contract", "Request.HistoryList", "Request.Contract.Supplier"})
	if history == nil {
		return nil, errs.ErrRequestHistoryNotFound
	}

	receiverList, _ := n.userRepo.FindAllByRoles([]string{constants.RolePlanner}, nil)
	for _, user := range []*entities.User{history.Request.Contractor, history.Request.Contract.Supplier} {
		if user.Id == history.CreatorId {
			continue
		}
		receiverList = append(receiverList, user)
	}

	var notificationList []*entities.Notification
	for _, receiver := range receiverList {
		if receiver.Id == history.CreatorId {
			continue
		}
		newNotification := entities.NewNotification(
			history.Status, false, constants.ObjTyRequest,
			history.RequestId, fmt.Sprintf("%v-%v", constants.ObjTyRequest, history.Request.Counter),
			history.CreatedAt, history.CreatorId, receiver.Id)
		newNotification.Object = history.Request
		notificationList = append(notificationList, newNotification)
	}

	return notificationList, nil
}

func (n NotificationFactory) CreateNotificationListForUserAction(senderId uuid.UUID, user *entities.User, action string) ([]*entities.Notification, error) {

	if user.Role == constants.RolePlanner {
		return nil, nil
	}

	receiverList, _ := n.userRepo.FindAllByRoles([]string{constants.RolePlanner}, nil)
	//receiverList = append(receiverList, user)

	var notificationList []*entities.Notification
	for _, receiver := range receiverList {
		if receiver.Id == user.Id {
			continue
		}
		notificationList = append(notificationList, entities.NewNotification(action, false, constants.ObjTyUser, user.Id, user.Email, time.Now(), senderId, receiver.Id))
	}

	return notificationList, nil
}
