package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u UserRepo) FindByEmail(email string, withs []string) (user *entities.User, err error) {

	err = u.db.Find(user, "users.email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
