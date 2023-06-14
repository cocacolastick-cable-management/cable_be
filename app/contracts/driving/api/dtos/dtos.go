package dtos

import "github.com/cable_management/cable_be/app/domain/services"

type CreateUserReq struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UpdateUserIsActiveReq struct {
	IsActive *bool `json:"isActive" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRes struct {
	Email     string              `json:"email"`
	Role      string              `json:"role"`
	Name      string              `json:"name"`
	AuthToken *services.AuthToken `json:"authToken"`
}

type CreateRequestReq struct {
	CableAmount     uint   `json:"cableAmount"`
	ContractCounter uint   `json:"contractCounter"`
	ContractorEmail string `json:"contractorEmail"`
}
