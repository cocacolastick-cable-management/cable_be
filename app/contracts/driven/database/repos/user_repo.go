package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
)

type IUserRepo interface {
	FindByEmail(email string, withs []string) (user *entities.User, err error)
	FindById(id uuid.UUID, withs []string) (user *entities.User, err error)
	Insert(user *entities.User) error
	Save(user *entities.User) error
	GetAll(withs []string) ([]*entities.User, error)
	FindAllByRoles(roles, withs []string) ([]*entities.User, error)
	FindAllActiveByRoles(roles, withs []string) ([]*entities.User, error)
}
