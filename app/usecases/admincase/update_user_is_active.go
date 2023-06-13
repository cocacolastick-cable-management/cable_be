package admincase

import (
	"errors"
	"github.com/cable_management/cable_be/app/contracts/database/repos"
	"github.com/cable_management/cable_be/app/contracts/email"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/app/usecases/_share/errs"
	"github.com/google/uuid"
)

var (
	ErrUserAlreadyDisable = errors.New("user is already disable")
	ErrUserAlreadyActive  = errors.New("user is already active")
)

type UpdateUserIsActiveReq struct {
	IsActive bool `json:"isActive" binding:"required"`
}

type IUpdateUserIsActive interface {
	Handle(accessToken string, userId uuid.UUID, req *UpdateUserIsActiveReq) error
}

type UpdateUserIsActive struct {
	userRepo      repos.IUserRepo
	authorService services.IAuthorizeService
	emailDriven   email.IEmail
}

func NewUpdateUserIsActive(userRepo repos.IUserRepo, authorService services.IAuthorizeService, emailDriven email.IEmail) *UpdateUserIsActive {
	return &UpdateUserIsActive{userRepo: userRepo, authorService: authorService, emailDriven: emailDriven}
}

func (u UpdateUserIsActive) Handle(accessToken string, userId uuid.UUID, req *UpdateUserIsActiveReq) error {

	// authorized
	_, err := u.authorService.Authorize(accessToken, []string{entities.RoleAdmin}, nil)
	if err != nil {
		return err
	}

	// find user
	user, _ := u.userRepo.FindById(userId, nil)
	if user == nil {
		return errs.ErrUserNotFound
	}

	// update user
	if user.IsActive == req.IsActive == true {
		return ErrUserAlreadyActive
	}
	if user.IsActive == req.IsActive == false {
		return ErrUserAlreadyDisable
	}
	user.IsActive = req.IsActive

	// save to database
	err = u.userRepo.Save(user)
	if err != nil {
		return err
	}

	// send email notif to current user
	go func() {
		err := u.emailDriven.SendEmailUpdateUserIsActive(email.EmailUpdateUserIsActiveDto{
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
