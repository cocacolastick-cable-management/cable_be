package plannercase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/go-playground/validator/v10"
)

type ICreateRequest interface {
	Handle(accessToken string, req *dtos.CreateRequestReq) error
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
}

func NewCreateRequest(
	authorService services.IAuthorizeService,
	validator *validator.Validate,
	requestFac services.IRequestFactory,
	requestRepo repos.IRequestRepo,
	requestHistoryFac services.IRequestHistoryFactory,
	mailFac services.IMailDataFactory,
	requestHistoryRepo repos.IRequestHistoryRepo,
	emailDriven email.IEmail) *CreateRequest {

	return &CreateRequest{
		authorService:      authorService,
		validator:          validator,
		requestFac:         requestFac,
		requestRepo:        requestRepo,
		requestHistoryFac:  requestHistoryFac,
		mailFac:            mailFac,
		requestHistoryRepo: requestHistoryRepo,
		emailDriven:        emailDriven}
}

func (c CreateRequest) Handle(accessToken string, req *dtos.CreateRequestReq) error {

	// authorize
	claims, err := c.authorService.Authorize(accessToken, []string{constants.RolePlanner}, nil)
	if err != nil {
		return err
	}

	// validate req
	err = c.validator.Struct(req)
	if err != nil {
		return err
	}

	// create req
	request, err := c.requestFac.CreateRequest(req.CableAmount, req.ContractCounter, req.ContractorEmail, claims.UserId)
	if err != nil {
		return err
	}
	history, err := c.requestHistoryFac.CreateRequestHistory(constants.ActionCreate, constants.StatusNew, request.Id, claims.UserId)
	if err != nil {
		return err
	}

	// insert database
	_ = c.requestRepo.Insert(request)
	_ = c.requestHistoryRepo.Insert(history)

	// send email notif
	go func() {
		mailList, _ := c.mailFac.CreateMailDataListForRequestAction(history.Id)
		_ = c.emailDriven.SendEmailOnRequestUpdate(mailList)
	}()

	// TODO sse

	return nil
}
