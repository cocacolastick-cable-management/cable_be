package dtos

type UpdateNotificationIsReadReq struct {
	IsRead *bool `json:"isRead" binding:"required"`
}
