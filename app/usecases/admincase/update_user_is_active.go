package admincase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/google/uuid"
)

type IUpdateUserIsActive interface {
	Handle(accessToken string, userId uuid.UUID, req *dtos.UpdateUserIsActiveReq) error
}

type UpdateUserIsActive struct {
	userRepo      repos.IUserRepo
	authorService services.IAuthorizeService
	emailDriven   email.IEmail
}

func NewUpdateUserIsActive(userRepo repos.IUserRepo, authorService services.IAuthorizeService, emailDriven email.IEmail) *UpdateUserIsActive {
	return &UpdateUserIsActive{userRepo: userRepo, authorService: authorService, emailDriven: emailDriven}
}

func (u UpdateUserIsActive) Handle(accessToken string, userId uuid.UUID, req *dtos.UpdateUserIsActiveReq) error {

	// authorized
	_, err := u.authorService.Authorize(accessToken, []string{constants.RoleAdmin}, nil)
	if err != nil {
		return err
	}

	// find user
	user, _ := u.userRepo.FindById(userId, nil)
	if user == nil {
		return errs.ErrUserNotFound
	}

	// update user
	if (user.IsActive == *req.IsActive) && (user.IsActive == true) {
		return errs.ErrUserAlreadyActive
	}
	if (user.IsActive == *req.IsActive) && (user.IsActive == false) {
		return errs.ErrUserAlreadyDisable
	}
	user.IsActive = *req.IsActive

	// save to database
	err = u.userRepo.Save(user)
	if err != nil {
		return err
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

	// send event notif to current user dashboard
	// send event notif to all planner
	// send email to relative users

	return nil
}
