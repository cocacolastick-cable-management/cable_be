package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RequestHistoryRepo struct {
	db *gorm.DB
}

func NewRequestHistoryRepo(db *gorm.DB) *RequestHistoryRepo {
	return &RequestHistoryRepo{db: db}
}

func (r RequestHistoryRepo) FindById(id uuid.UUID, withs []string) (*entities.RequestHistory, error) {

	query := r.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	history := new(entities.RequestHistory)
	result := query.First(&history, "request_histories.id = ?", id)

	return history, result.Error
}

func (r RequestHistoryRepo) Insert(history *entities.RequestHistory) error {
	result := r.db.Create(history)
	return result.Error
}
