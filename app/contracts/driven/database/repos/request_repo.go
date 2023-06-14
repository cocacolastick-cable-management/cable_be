package repos

import "github.com/cable_management/cable_be/app/domain/entities"

type IRequestRepo interface {
	Insert(request *entities.Request) error
}
