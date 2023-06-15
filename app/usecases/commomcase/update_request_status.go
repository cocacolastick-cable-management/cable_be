package commomcase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/elliotchance/pie/v2"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IUpdateRequestStatus interface {
	Handle(accessToken string, requestId uuid.UUID, req *dtos.UpdateRequestStatusRequest) error
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
}

func NewUpdateRequestStatus(
	authorService services.IAuthorizeService,
	requestRepo repos.IRequestRepo,
	requestHistoryRepo repos.IRequestHistoryRepo,
	userRepo repos.IUserRepo,
	validator *validator.Validate,
	requestHistoryFac services.IRequestHistoryFactory,
	mailDataFactory services.IMailDataFactory,
	emailDriven email.IEmail) *UpdateRequestStatus {

	return &UpdateRequestStatus{
		authorService:      authorService,
		requestRepo:        requestRepo,
		requestHistoryRepo: requestHistoryRepo,
		userRepo:           userRepo,
		validator:          validator,
		requestHistoryFac:  requestHistoryFac,
		mailDataFactory:    mailDataFactory,
		emailDriven:        emailDriven}
}

func (u UpdateRequestStatus) Handle(accessToken string, requestId uuid.UUID, req *dtos.UpdateRequestStatusRequest) error {

	// authorize
	claims, err := u.authorService.Authorize(accessToken, []string{constants.RolePlanner, constants.RoleContractor, constants.RoleSupplier}, nil)
	if err != nil {
		return err
	}

	// validate req.Status
	if !pie.Contains(constants.StatusList, req.Status) {
		return errs.ErrInvalidRequestStatus
	}

	// check permission
	action := ""
	switch {
	case claims.Role == constants.RolePlanner && req.Status == constants.StatusCanceled:
		action = constants.ActionCancel
	case claims.Role == constants.RoleSupplier && req.Status == constants.StatusReady:
		action = constants.ActionUpdate
	case claims.Role == constants.RoleContractor && req.Status == constants.StatusCollected:
		action = constants.ActionUpdate
	default:
		return errs.ErrUnauthorized
	}

	// update request
	request, err := u.requestRepo.FindById(requestId, nil)
	if err != nil {
		return errs.ErrRequestNotFound
	}
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

	// TODO sse

	return nil
}
