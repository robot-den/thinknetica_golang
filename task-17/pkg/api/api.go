// package api реализует HTTP API для сервиса авторизации
package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"pkg/auth"
	"pkg/model"
)

const (
	AuthPath = "/auth"
	RootPath = "/"
)

// API представляет собой объект реализующий методы-обработчики
type API struct {
	authorizator *auth.Auth
	port         string
	router       *mux.Router
}

// New создает объект API
func New(authorizator *auth.Auth, port string) *API {
	s := API{
		authorizator: authorizator,
		port:         port,
		router:       mux.NewRouter(),
	}
	return &s
}

// Serve подготавливает и запускает HTTP сервер
func (a *API) Serve() {
	a.endpoints()
	loggedRouter := handlers.LoggingHandler(os.Stdout, a.router)

	err := http.ListenAndServe(a.port, loggedRouter)
	if err != nil {
		fmt.Println(err)
	}
}

// endpoints регистрирует обработчики
func (a *API) endpoints() {
	a.router.HandleFunc(AuthPath, a.handleAuth).Methods(http.MethodPost)
	a.router.HandleFunc(RootPath, a.handleOptions).Methods(http.MethodOptions)
}

// handleAuth обрабатывает запросы содержащие логин/пароль и возвращает токен с правами доступа
func (a *API) handleAuth(w http.ResponseWriter, r *http.Request) {
	creds := model.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	token, err := a.authorizator.Authorize(creds)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write([]byte(token))
	if err != nil {
		fmt.Println(err)
	}
}

func (a *API) handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "OPTIONS, POST")
}
