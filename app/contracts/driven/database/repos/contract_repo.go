package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
)

type IContractRepo interface {
	FindByCounter(counter uint, withs []string) (*entities.Contract, error)
	GetAll(withs []string) ([]*entities.Contract, error)
	GetBySupplierId(supplierId uuid.UUID, withs []string) ([]*entities.Contract, error)
}
