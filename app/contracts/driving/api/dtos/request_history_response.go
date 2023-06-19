package dtos

import (
	"github.com/google/uuid"
	"time"
)

type RequestHistoryRes struct {
	Id           uuid.UUID `json:"id"`
	RequestId    uuid.UUID `json:"requestId"`
	CreatorId    uuid.UUID `json:"creatorId"`
	CreatorEmail string    `json:"creatorEmail"`
	CreatorRole  string    `json:"creatorRole"`
	CreatedAt    time.Time `json:"createdAt"`
	Status       string    `json:"status"`
	Action       string    `json:"action"`
}
