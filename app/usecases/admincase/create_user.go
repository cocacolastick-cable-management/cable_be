package admincase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driven/sse"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/go-playground/validator/v10"
)

type ICreateUser interface {
	Handle(accessToken string, req *dtos.CreateUserReq) (userRes *dtos.UserRes, err error)
}

type CreateUser struct {
	userRepo         repos.IUserRepo
	userFac          services.IUserFactory
	validator        *validator.Validate
	authorService    services.IAuthorizeService
	emailDriven      email.IEmail
	passwordService  services.IPasswordService
	notificationFac  services.INotificationFactory
	notificationRepo repos.INotificationRepo
	sseDriven        sse.ISSEDriven
}

func NewCreateUser(
	userRepo repos.IUserRepo,
	userFac services.IUserFactory,
	validator *validator.Validate,
	authorService services.IAuthorizeService,
	emailDriven email.IEmail,
	passwordService services.IPasswordService,
	notificationFac services.INotificationFactory,
	notificationRepo repos.INotificationRepo,
	sseDriven sse.ISSEDriven) *CreateUser {

	return &CreateUser{
		userRepo:         userRepo,
		userFac:          userFac,
		validator:        validator,
		authorService:    authorService,
		emailDriven:      emailDriven,
		passwordService:  passwordService,
		notificationFac:  notificationFac,
		notificationRepo: notificationRepo,
		sseDriven:        sseDriven}
}

func (c CreateUser) Handle(accessToken string, req *dtos.CreateUserReq) (userRes *dtos.UserRes, err error) {

	// authorize
	claims, err := c.authorService.Authorize(accessToken, []string{constants.RoleAdmin}, nil)
	if err != nil {
		return nil, err
	}

	// validate
	err = c.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	// create user
	password := c.passwordService.GeneratePassword(10)
	newUser, err := c.userFac.CreateUser(req.Role, req.Email, req.Name, password)
	if err != nil {
		return nil, err
	}

	// insert to database
	err = c.userRepo.Insert(newUser)
	if err != nil {
		return nil, err
	}

	// send email with account to user
	go func() {
		_ = c.emailDriven.SendEmailNewUser(email.ToMailNewUserDto(newUser, password))
	}()

	userRes, _ = maps.ToUserRes(newUser)

	// create and send notification
	go func() {

		notificationList, _ := c.notificationFac.CreateNotificationListForUserAction(claims.UserId, newUser, constants.ActionCreate)

		go func() {
			_ = c.notificationRepo.InsertMany(notificationList)
		}()

		go func() {
			notificationDtoList := make([]*sse.Message, len(notificationList))
			for i, notification := range notificationList {
				notificationDtoList[i], _ = sse.ToMessage(notification.ReceiverId, claims.UserEmail, notification, userRes)
			}
			_ = c.sseDriven.SendMessage(notificationDtoList)
		}()
	}()

	return userRes, nil
}

// validation
