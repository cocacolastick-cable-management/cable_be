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

func (r RequestRepo) Save(request *entities.Request) error {
	result := r.db.Save(request)
	return result.Error
}

func (r RequestRepo) FindById(id uuid.UUID, withs []string) (*entities.Request, error) {

	query := r.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	var request *entities.Request
	result := query.First(&request, "requests.id = ?", id)

	return request, result.Error
}

func (r RequestRepo) GetAll(withs []string) ([]*entities.Request, error) {

	var RequestList []*entities.Request
	query := r.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.Find(&RequestList)
	if result.Error != nil {
		return nil, result.Error
	}

	return RequestList, nil
}
