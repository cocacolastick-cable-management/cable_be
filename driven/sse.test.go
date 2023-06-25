package driven

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan ClientChan

	//// Closed client connections
	//ClosedClients chan ClientChan
	//
	//// Total client connections
	TotalClients map[ClientChan]bool
}

type ClientChan struct {
	Message chan string
}

func NewSSEServer() (event *Event) {
	event = &Event{
		Message:    make(chan string),
		NewClients: make(chan ClientChan),
		//ClosedClients: make(chan ClientChan),
		TotalClients: make(map[ClientChan]bool),
	}

	go event.listen()

	return event
}

func (stream *Event) listen() {
	for {
		select {
		// Add new available client
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		//// Remove closed client
		//case client := <-stream.ClosedClients:
		//	delete(stream.TotalClients, client)
		//	//close(client)
		//	log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan.Message <- eventMsg
			}
		}
	}
}

func (stream *Event) ServeHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := ClientChan{
			Message: make(chan string),
		}

		// Send new connection to event server
		stream.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			//stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}

//func BuildSSETest(engine *gin.Engine) {
//
//	sseStream := sse.NewSSEServer()
//
//	go func() {
//		for {
//			time.Sleep(time.Second * 5)
//			now := time.Now().Format("2006-01-02 15:04:05")
//			currentTime := fmt.Sprintf("The Current Time Is %v", now)
//
//			// Send current time to clients message channel
//			sseStream.Message <- currentTime
//		}
//	}()
//
//	engine.GET("/sse", sse.HeadersMiddleware(), sseStream.ServeHTTP(), func(c *gin.Context) {
//		v, ok := c.Get("clientChan")
//		if !ok {
//			return
//		}
//		clientChan, ok := v.(sse.clientChan)
//		if !ok {
//			return
//		}
//		c.Stream(func(w io.Writer) bool {
//			// Stream message to client from message channel
//			if msg, ok := <-clientChan; ok {
//				c.SSEvent("message", struct {
//					Yeah string
//					Msg  any
//				}{
//					Yeah: "whats up bitches",
//					Msg:  msg,
//				})
//				return true
//			}
//			return false
//		})
//	})
//}
