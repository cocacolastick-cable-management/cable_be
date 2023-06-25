package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u UserRepo) FindByEmail(email string, withs []string) (user *entities.User, err error) {

	result := u.db.First(&user, "users.email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u UserRepo) FindById(id uuid.UUID, withs []string) (user *entities.User, err error) {

	err = u.db.First(&user, "users.id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserRepo) Insert(user *entities.User) error {
	result := u.db.Create(user)
	return result.Error
}

func (u UserRepo) Save(user *entities.User) error {
	result := u.db.Save(user)
	return result.Error
}

func (u UserRepo) GetAll(withs []string) ([]*entities.User, error) {

	var userList []*entities.User
	query := u.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.Order("users.created_at desc").Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}

	return userList, nil
}

func (u UserRepo) FindAllByRoles(roles, withs []string) ([]*entities.User, error) {

	var userList []*entities.User
	query := u.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.Order("users.created_at desc").Find(&userList, "users.role IN ?", roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return userList, nil
}

func (u UserRepo) FindAllActiveByRoles(roles, withs []string) ([]*entities.User, error) {

	var userList []*entities.User
	query := u.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.Order("users.created_at desc").Find(&userList, "users.is_active = ? AND users.role IN ?", true, roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return userList, nil
}
