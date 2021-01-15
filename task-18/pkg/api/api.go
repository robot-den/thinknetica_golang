// package api реализует HTTP API для чат-сервера
package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"task-18/pkg/chat"
)

const (
	MessagePath  = "/send"
	MessagesPath = "/messages"
)

// API представляет собой объект реализующий методы-обработчики
type API struct {
	port   string
	router *mux.Router
	chat   *chat.Chat
}

// New создает объект API
func New(chat *chat.Chat, port string) *API {
	s := API{
		port:   port,
		router: mux.NewRouter(),
		chat:   chat,
	}
	return &s
}

// Serve подготавливает и запускает HTTP сервер
func (a *API) Serve() {
	a.endpoints()

	err := http.ListenAndServe(a.port, a.router)
	if err != nil {
		fmt.Println(err)
	}
}

// endpoints регистрирует обработчики
func (a *API) endpoints() {
	a.router.HandleFunc(MessagePath, a.handleMessage).Methods(http.MethodGet)
	a.router.HandleFunc(MessagesPath, a.handleMessages).Methods(http.MethodGet)
}
