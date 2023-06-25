package admincase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driven/sse"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/google/uuid"
)

type IUpdateUserIsActive interface {
	Handle(accessToken string, userId uuid.UUID, req *dtos.UpdateUserIsActiveReq) (*dtos.UserRes, error)
}

type UpdateUserIsActive struct {
	userRepo         repos.IUserRepo
	authorService    services.IAuthorizeService
	emailDriven      email.IEmail
	notificationFac  services.INotificationFactory
	notificationRepo repos.INotificationRepo
	sseDriven        sse.ISSEDriven
}

func NewUpdateUserIsActive(userRepo repos.IUserRepo, authorService services.IAuthorizeService, emailDriven email.IEmail, notificationFac services.INotificationFactory, notificationRepo repos.INotificationRepo, sseDriven sse.ISSEDriven) *UpdateUserIsActive {
	return &UpdateUserIsActive{userRepo: userRepo, authorService: authorService, emailDriven: emailDriven, notificationFac: notificationFac, notificationRepo: notificationRepo, sseDriven: sseDriven}
}

func (u UpdateUserIsActive) Handle(accessToken string, userId uuid.UUID, req *dtos.UpdateUserIsActiveReq) (*dtos.UserRes, error) {

	// authorized
	claims, err := u.authorService.Authorize(accessToken, []string{constants.RoleAdmin}, nil)
	if err != nil {
		return nil, err
	}

	// find user
	user, _ := u.userRepo.FindById(userId, nil)
	if user == nil {
		return nil, errs.ErrUserNotFound
	}

	// update user
	if (user.IsActive == *req.IsActive) && (user.IsActive == true) {
		return nil, errs.ErrUserAlreadyActive
	}
	if (user.IsActive == *req.IsActive) && (user.IsActive == false) {
		return nil, errs.ErrUserAlreadyDisable
	}
	user.IsActive = *req.IsActive

	// save to database
	err = u.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	// send email notif to current user
	go func() {
		err := u.emailDriven.SendEmailUpdateUserIsActive(email.MailUpdateUserIsActiveDto{
			NewStatus: user.IsActive,
			Email:     user.Email,
		})
		if err != nil {
			// TODO logger
		}
	}()

	userRes, _ := maps.ToUserRes(user)

	// create and send notification
	go func() {

		action := constants.ActionEnable
		if !*req.IsActive {
			action = constants.ActionDisable
		}

		notificationList, _ := u.notificationFac.CreateNotificationListForUserAction(claims.UserId, user, action)

		go func() {
			_ = u.notificationRepo.InsertMany(notificationList)
		}()

		go func() {
			notificationDtoList := make([]*sse.Message, len(notificationList))
			for i, notification := range notificationList {
				notificationDtoList[i], _ = sse.ToMessage(notification.ReceiverId, claims.UserEmail, notification, userRes)
			}
			_ = u.sseDriven.SendMessage(notificationDtoList)
		}()
	}()

	return userRes, nil
}
