package services

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/elliotchance/pie/v2"
	"github.com/google/uuid"
	"time"
)

type IRequestHistoryFactory interface {
	CreateRequestHistory(action, status string, requestId, creatorId uuid.UUID) (*entities.RequestHistory, error)
}

type RequestHistoryFactory struct {
	requestRepo repos.IRequestRepo
}

func NewRequestHistoryFactory(requestRepo repos.IRequestRepo) *RequestHistoryFactory {
	return &RequestHistoryFactory{requestRepo: requestRepo}
}

func (r RequestHistoryFactory) CreateRequestHistory(action, status string, requestId, creatorId uuid.UUID) (*entities.RequestHistory, error) {

	// validate action
	if !pie.Contains(constants.ActionList, action) {
		return nil, errs.ErrInvalidRequestAction
	}

	// validate status
	if !pie.Contains(constants.StatusList, status) {
		return nil, errs.ErrInvalidRequestStatus
	}

	// TODO validate requestId
	//request, _ := r.requestRepo.FindById(requestId, []string{""})

	// TODO validate creatorId

	// TODO create history

	return entities.NewRequestHistory(action, status, creatorId, requestId, time.Now()), nil
}
