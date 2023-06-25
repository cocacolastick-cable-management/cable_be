package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
)

type IRequestRepo interface {
	Insert(request *entities.Request) error
	Save(request *entities.Request) error
	FindById(id uuid.UUID, withs []string) (*entities.Request, error)
	GetAll(withs []string) ([]*entities.Request, error)
	GetBySupplierId(supplierId uuid.UUID, withs []string) ([]*entities.Request, error)
	GetByContractorId(supplierId uuid.UUID, withs []string) ([]*entities.Request, error)
}
