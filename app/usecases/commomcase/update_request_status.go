package commomcase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driven/sse"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/maps"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/elliotchance/pie/v2"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IUpdateRequestStatus interface {
	Handle(accessToken string, requestId uuid.UUID, req *dtos.UpdateRequestStatusRequest) (*dtos.RequestRes, error)
}

type UpdateRequestStatus struct {
	authorService      services.IAuthorizeService
	requestRepo        repos.IRequestRepo
	requestHistoryRepo repos.IRequestHistoryRepo
	userRepo           repos.IUserRepo
	validator          *validator.Validate
	requestHistoryFac  services.IRequestHistoryFactory
	mailDataFactory    services.IMailDataFactory
	emailDriven        email.IEmail
	notificationFac    services.INotificationFactory
	notificationRepo   repos.INotificationRepo
	sseDriven          sse.ISSEDriven
}

func NewUpdateRequestStatus(
	authorService services.IAuthorizeService,
	requestRepo repos.IRequestRepo,
	requestHistoryRepo repos.IRequestHistoryRepo,
	userRepo repos.IUserRepo,
	validator *validator.Validate,
	requestHistoryFac services.IRequestHistoryFactory,
	mailDataFactory services.IMailDataFactory,
	emailDriven email.IEmail,
	notificationFac services.INotificationFactory,
	notificationRepo repos.INotificationRepo,
	sseDriven sse.ISSEDriven) *UpdateRequestStatus {

	return &UpdateRequestStatus{
		authorService:      authorService,
		requestRepo:        requestRepo,
		requestHistoryRepo: requestHistoryRepo,
		userRepo:           userRepo,
		validator:          validator,
		requestHistoryFac:  requestHistoryFac,
		mailDataFactory:    mailDataFactory,
		emailDriven:        emailDriven,
		notificationFac:    notificationFac,
		notificationRepo:   notificationRepo,
		sseDriven:          sseDriven}
}

func (u UpdateRequestStatus) Handle(accessToken string, requestId uuid.UUID, req *dtos.UpdateRequestStatusRequest) (*dtos.RequestRes, error) {

	// authorize
	claims, err := u.authorService.Authorize(accessToken, []string{constants.RolePlanner, constants.RoleContractor, constants.RoleSupplier}, nil)
	if err != nil {
		return nil, err
	}

	// validate req.Status
	if !pie.Contains(constants.StatusList, req.Status) {
		return nil, errs.ErrInvalidRequestStatus
	}

	// find request
	request, err := u.requestRepo.FindById(requestId, nil)
	if err != nil {
		return nil, errs.ErrRequestNotFound
	}

	// check permission
	action := ""
	switch {
	case claims.Role == constants.RolePlanner && req.Status == constants.StatusCanceled && request.Status != constants.StatusCanceled:
		action = constants.ActionCancel
	case claims.Role == constants.RoleSupplier && req.Status == constants.StatusReady && request.Status == constants.StatusNew:
		action = constants.ActionUpdate
	case claims.Role == constants.RoleContractor && req.Status == constants.StatusCollected && request.Status == constants.StatusReady:
		action = constants.ActionUpdate
	default:
		return nil, errs.ErrUnauthorized
	}

	// update request
	request.Status = req.Status

	// create history
	history, _ := u.requestHistoryFac.CreateRequestHistory(action, req.Status, request.Id, claims.UserId)

	// save to db
	_ = u.requestRepo.Save(request)
	_ = u.requestHistoryRepo.Insert(history)

	// send email noti
	go func() {
		mailList, _ := u.mailDataFactory.CreateMailDataListForRequestAction(history.Id)
		_ = u.emailDriven.SendEmailOnRequestUpdate(mailList)
	}()

	request, _ = u.requestRepo.FindById(request.Id, []string{"Contractor", "Contract", "Contract.Supplier", "Planner", "HistoryList", "HistoryList.Creator"})
	requestRes, _ := maps.ToRequestResponse(request)

	go func() {
		notificationList, _ := u.notificationFac.CreateNotificationListForRequestAction(history.Id)
		go func() {
			_ = u.notificationRepo.InsertMany(notificationList)
		}()
		go func() {
			notificationDtoList := make([]*sse.Message, len(notificationList))
			for i, notification := range notificationList {
				notificationDtoList[i], _ = sse.ToMessage(notification.ReceiverId, claims.UserEmail, notification, requestRes)
			}
			_ = u.sseDriven.SendMessage(notificationDtoList)
		}()
	}()

	return requestRes, nil
}
