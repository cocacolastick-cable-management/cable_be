package dtos

import "github.com/cable_management/cable_be/app/domain/services"

type AuthRes struct {
	Email     string              `json:"email"`
	Role      string              `json:"role"`
	Name      string              `json:"name"`
	AuthToken *services.AuthToken `json:"authToken"`
}
