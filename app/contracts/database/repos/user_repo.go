package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
)

type IUserRepo interface {
	FindByEmail(email string, withs []string) (user *entities.User, err error)
	Insert(user *entities.User) error
}
