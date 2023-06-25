package plannercase

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

type ICreateRequest interface {
	Handle(accessToken string, req *dtos.CreateRequestReq) (*dtos.RequestRes, error)
}

type CreateRequest struct {
	authorService      services.IAuthorizeService
	validator          *validator.Validate
	requestFac         services.IRequestFactory
	requestRepo        repos.IRequestRepo
	requestHistoryFac  services.IRequestHistoryFactory
	mailFac            services.IMailDataFactory
	requestHistoryRepo repos.IRequestHistoryRepo
	emailDriven        email.IEmail
	notificationFac    services.INotificationFactory
	notificationRepo   repos.INotificationRepo
	sseDriven          sse.ISSEDriven
}

func NewCreateRequest(
	authorService services.IAuthorizeService,
	validator *validator.Validate,
	requestFac services.IRequestFactory,
	requestRepo repos.IRequestRepo,
	requestHistoryFac services.IRequestHistoryFactory,
	mailFac services.IMailDataFactory,
	requestHistoryRepo repos.IRequestHistoryRepo,
	emailDriven email.IEmail,
	notificationFac services.INotificationFactory,
	notificationRepo repos.INotificationRepo,
	sseDriven sse.ISSEDriven) *CreateRequest {

	return &CreateRequest{
		authorService:      authorService,
		validator:          validator,
		requestFac:         requestFac,
		requestRepo:        requestRepo,
		requestHistoryFac:  requestHistoryFac,
		mailFac:            mailFac,
		requestHistoryRepo: requestHistoryRepo,
		emailDriven:        emailDriven,
		notificationFac:    notificationFac,
		notificationRepo:   notificationRepo,
		sseDriven:          sseDriven}
}

func (c CreateRequest) Handle(accessToken string, req *dtos.CreateRequestReq) (*dtos.RequestRes, error) {

	// authorize
	claims, err := c.authorService.Authorize(accessToken, []string{constants.RolePlanner}, nil)
	if err != nil {
		return nil, err
	}

	// validate req
	err = c.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	// create req
	request, err := c.requestFac.CreateRequest(req.CableAmount, req.ContractCounter, req.ContractorEmail, claims.UserId)
	if err != nil {
		return nil, err
	}
	history, err := c.requestHistoryFac.CreateRequestHistory(constants.ActionCreate, constants.StatusNew, request.Id, claims.UserId)
	if err != nil {
		return nil, err
	}

	// insert database
	_ = c.requestRepo.Insert(request)
	_ = c.requestHistoryRepo.Insert(history)

	// send email notif
	go func() {
		mailList, _ := c.mailFac.CreateMailDataListForRequestAction(history.Id)
		_ = c.emailDriven.SendEmailOnRequestUpdate(mailList)
	}()

	// query to create response to return client
	request, _ = c.requestRepo.FindById(request.Id, []string{"Contractor", "Contract", "Contract.Supplier", "Planner", "HistoryList", "HistoryList.Creator"})
	requestRes, _ := maps.ToRequestResponse(request)

	// notification
	go func() {
		notificationList, _ := c.notificationFac.CreateNotificationListForRequestAction(history.Id)
		go func() {
			_ = c.notificationRepo.InsertMany(notificationList)
		}()
		go func() {
			notificationDtoList := make([]*sse.Message, len(notificationList))
			for i, notification := range notificationList {
				notificationDtoList[i], _ = sse.ToMessage(notification.ReceiverId, claims.UserEmail, notification, requestRes)
			}
			_ = c.sseDriven.SendMessage(notificationDtoList)
		}()
	}()

	return requestRes, nil
}
