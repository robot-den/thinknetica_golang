package main

import (
	"task-18/pkg/api"
	"task-18/pkg/chat"
)

// Service представляет собой чат-сервер
type Service struct {
	api  *api.API
	chat *chat.Chat
}

func main() {
	service := new()

	service.api.Serve()
}

func new() *Service {
	chat := chat.New()
	s := Service{
		api:  api.New(chat, ":9000"),
		chat: chat,
	}

	return &s
}
