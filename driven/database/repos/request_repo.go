package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
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

func (r RequestRepo) FindById(id uuid.UUID, withs []string) (*entities.Request, error) {

	query := r.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	request := new(entities.Request)
	result := query.First(&request, "request.id = ?", id)

	return request, result.Error
}
