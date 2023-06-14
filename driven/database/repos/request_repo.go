package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"gorm.io/gorm"
)

type RequestRepo struct {
	db *gorm.DB
}

func NewRequestRepo(db *gorm.DB) *RequestRepo {
	return &RequestRepo{db: db}
}

func (r RequestRepo) Insert(request *entities.Request) error {
	result := r.db.Create(request)
	return result.Error
}
