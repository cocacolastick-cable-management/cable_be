package services

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/cable_management/cable_be/app/domain/errs"
	"github.com/google/uuid"
)

type IMailDataFactory interface {
	CreateMailDataListForRequestAction(historyId uuid.UUID) ([]email.MailRequestActionDto, error)
}

type MailDataFactory struct {
	userRepo           repos.IUserRepo
	requestHistoryRepo repos.IRequestHistoryRepo
}

func NewMailDataFactory(userRepo repos.IUserRepo, requestHistoryRepo repos.IRequestHistoryRepo) *MailDataFactory {
	return &MailDataFactory{userRepo: userRepo, requestHistoryRepo: requestHistoryRepo}
}

func (e MailDataFactory) CreateMailDataListForRequestAction(historyId uuid.UUID) ([]email.MailRequestActionDto, error) {

	// find history
	history, _ := e.requestHistoryRepo.FindById(historyId,
		[]string{"Creator", "Request", "Request.Planner", "Request.Contractor", "Request.Contract", "Request.Contract.Supplier"})
	if history == nil {
		return nil, errs.ErrRequestHistoryNotFound
	}

	// check if any related user is disabled
	//switch {
	//case !history.Request.Contractor.IsActive:
	//	return nil, errs.ErrContractorIsDisable
	//case !history.Request.Contract.Supplier.IsActive:
	//	return nil, errs.ErrSupplierIsDisable
	//	//case !history.Request.Planner.IsActive:
	//	//	return nil, errs.ErrPlannerIsDisable
	//}

	// get related users
	relatedUserList := make([]*entities.User, 0, 3)
	relatedUserList = append(
		relatedUserList,
		history.Request.Contractor,
		history.Request.Planner,
		history.Request.Contract.Supplier)

	// create mail list
	var mailList []email.MailRequestActionDto

	for _, user := range relatedUserList {
		if user.Id == history.CreatorId {
			continue
		}
		mailList = append(mailList, email.MailRequestActionDto{
			SenderEmail:     history.Creator.Email,
			ReceiverEmail:   user.Email,
			Action:          history.Action,
			RequestCounter:  history.Request.Counter,
			ContractCounter: history.Request.Contract.Counter,
			Status:          history.Status,
			Time:            history.CreatedAt,
		})
	}

	return mailList, nil
}
