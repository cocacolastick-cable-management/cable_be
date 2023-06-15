package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
)

type IRequestHistoryRepo interface {
	FindById(id uuid.UUID, withs []string) (*entities.RequestHistory, error)
	Insert(history *entities.RequestHistory) error
}
