package maps

import (
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/entities"
)

func ToUserRes(user *entities.User) (*dtos.UserRes, error) {
	return &dtos.UserRes{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}, nil
}
