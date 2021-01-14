package main

import (
	"pkg/api"
	"pkg/auth"
)

// Service представляет собой сервер авторизации
type Service struct {
	auth *auth.Auth
	api  *api.API
}

func main() {
	service := new()

	service.api.Serve()
}

func new() *Service {
	auth := auth.New()

	s := Service{
		auth: auth,
		api:  api.New(auth, ":9000"),
	}

	return &s
}
