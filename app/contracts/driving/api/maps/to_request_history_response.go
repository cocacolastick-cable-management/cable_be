package maps

import (
	"github.com/cable_management/cable_be/_share/errs"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/entities"
)

func ToRequestHistoryResponse(history *entities.RequestHistory) (*dtos.RequestHistoryRes, error) {

	if history.Creator == nil {
		return nil, errs.ErrNullReference
	}

	return &dtos.RequestHistoryRes{
		Id:           history.Id,
		RequestId:    history.RequestId,
		CreatorId:    history.CreatorId,
		CreatorEmail: history.Creator.Email,
		CreatorRole:  history.Creator.Role,
		CreatedAt:    history.CreatedAt,
		Status:       history.Status,
		Action:       history.Action,
	}, nil
}

//type RequestHistoryRes struct {
//	Id           uuid.UUID
//	RequestId    uuid.UUID
//	CreatorId    uuid.UUID
//	CreatorName  string
//	CreatorEmail string
//	CreatorRole  string
//	CreatedAt    time.Time
//	Status       string
//	Action       string
//}
