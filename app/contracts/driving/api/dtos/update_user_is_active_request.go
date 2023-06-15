package dtos

type UpdateUserIsActiveReq struct {
	IsActive *bool `json:"isActive" binding:"required"`
}
