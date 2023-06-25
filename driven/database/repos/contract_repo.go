package repos

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContractRepo struct {
	db *gorm.DB
}

func NewContractRepo(db *gorm.DB) *ContractRepo {
	return &ContractRepo{db: db}
}

func (c ContractRepo) FindByCounter(counter uint, withs []string) (*entities.Contract, error) {

	query := c.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	contract := new(entities.Contract)
	result := query.First(&contract, "contracts.counter = ?", counter)

	return contract, result.Error
}

func (c ContractRepo) GetAll(withs []string) ([]*entities.Contract, error) {

	var contractList []*entities.Contract
	query := c.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.Order("contracts.created_at desc").Find(&contractList)
	if result.Error != nil {
		return nil, result.Error
	}

	return contractList, nil
}

func (c ContractRepo) GetBySupplierId(supplierId uuid.UUID, withs []string) ([]*entities.Contract, error) {

	var contractList []*entities.Contract
	query := c.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.
		Order("contracts.created_at desc").
		Find(&contractList, "contracts.supplier_id = ?", supplierId)
	if result.Error != nil {
		return nil, result.Error
	}

	return contractList, nil
}
