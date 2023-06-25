package sse

import (
	"github.com/cable_management/cable_be/app/contracts/driven/sse"
)

//type NotificationChanType chan sse.Message
//
//var NotificationChan = make(NotificationChanType)

type SSEDriven struct {
	sseServer ISSEServer
}

func NewSSEDriven(sseServer ISSEServer) *SSEDriven {
	return &SSEDriven{sseServer: sseServer}
}

func (s SSEDriven) SendMessage(notificationList []*sse.Message) error {

	for _, notification := range notificationList {
		s.sseServer.GetEventChan() <- Event{
			ReceiverId: notification.ReceiverId,
			Data:       notification,
		}
	}

	return nil
}
