package repos

import "github.com/cable_management/cable_be/app/domain/entities"

type IContractRepo interface {
	FindByCounter(counter uint, withs []string) (*entities.Contract, error)
}
