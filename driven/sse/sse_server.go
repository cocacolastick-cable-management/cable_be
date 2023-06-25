package sse

import (
	"fmt"
	"github.com/cable_management/cable_be/_share/env"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

type Client struct {
	Id        uuid.UUID
	EventChan chan any
}

type Event struct {
	ReceiverId uuid.UUID `json:"receiverId"`
	Data       any       `json:"data"`
}

type ISSEServer interface {
	Register(c gin.IRouter)
	GetEventChan() chan Event
	listen()
	serveHTTP(c *gin.Context)
	handleSendEvent(c *gin.Context)
}

type SSEServer struct {
	eventChan     chan Event
	newClientChan chan Client
	closedClient  chan Client
	clientChan    map[Client]bool
	env           env.Env
}

func NewSSEServer(env env.Env) (server *SSEServer) {

	server = &SSEServer{
		eventChan:     make(chan Event),
		newClientChan: make(chan Client),
		closedClient:  make(chan Client),
		clientChan:    make(map[Client]bool),
		env:           env,
	}

	go server.listen()

	return server
}

func (s *SSEServer) GetEventChan() chan Event {
	return s.eventChan
}

func (s *SSEServer) listen() {

	for {

		select {

		case client := <-s.newClientChan:
			s.clientChan[client] = true
			log.Printf("Client %v added. %d registered clients", client.Id, len(s.clientChan))

		case client := <-s.closedClient:
			delete(s.clientChan, client)
			close(client.EventChan)
			log.Printf("Removed client %v. %d registered clients", client.Id, len(s.clientChan))

		case event := <-s.eventChan:
			for client := range s.clientChan {
				if client.Id == event.ReceiverId {
					client.EventChan <- event
				}
			}
		}
	}
}

const (
	ClientKey = "client"
)

func (s *SSEServer) serveHTTP(c *gin.Context) {

	accessToken := c.Query("access_token")

	claims := &services.AuthTokenClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.env.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, middlewares.Unauthenticated)
		return
	}

	client := Client{
		Id:        claims.UserId,
		EventChan: make(chan any),
	}

	s.newClientChan <- client

	defer func() {
		s.closedClient <- client
	}()

	c.Set(ClientKey, client)

	c.Next()
}

func headersMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

	c.Next()
}

func (s *SSEServer) Register(c gin.IRouter) {
	sseRoute := c.Group("/sse")

	sseRoute.GET("/notifications",
		headersMiddleware,
		//middlewares.ParseAccessToken, // TODO, really bad code, currently have no idea to fix
		s.serveHTTP,
		s.handleSendEvent,
		middlewares.HandleGlobalErrors, // TODO, really bad code, currently have no idea to fix
	)
}

func (s *SSEServer) handleSendEvent(c *gin.Context) {

	clientRaw := c.MustGet(ClientKey)

	client := clientRaw.(Client)

	w := c.Writer
	if flusher, ok := w.(http.Flusher); ok {
		_, _ = fmt.Fprintf(w, "connect success")
		flusher.Flush()
	}

	c.Stream(func(w io.Writer) bool {
		if event, ok := <-client.EventChan; ok {
			c.SSEvent("message", event)
			return true
		}
		return false
	})
}
