// package api реализует HTTP API поисковика
package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"pkg/engine"
	"pkg/model"
	"strconv"
)

const (
	DocumentsPath       = "/documents"
	DocumentsSearchPath = "/documents/search"
	RootPath            = "/"
)

// Service представляет собой объект реализующий методы-обработчики API
type Service struct {
	engine *engine.Service
	port   string
	router *mux.Router
}

// NewService создает объект сервиса
func NewService(eng *engine.Service, port string) *Service {
	s := Service{
		engine: eng,
		port:   port,
		router: mux.NewRouter(),
	}
	return &s
}

// Serve подготавливает и запускает HTTP сервер
func (s *Service) Serve() {
	s.endpoints()
	err := http.ListenAndServe(s.port, s.router)
	if err != nil {
		fmt.Println(err)
	}
}

// endpoints регистрирует обработчики и возвращает роутер
func (s *Service) endpoints() {
	s.router.HandleFunc(DocumentsSearchPath, s.handleSearch).Methods(http.MethodGet)
	s.router.HandleFunc(DocumentsPath, s.handleCreateDocument).Methods(http.MethodPost)
	s.router.HandleFunc(DocumentsPath+"/{id}", s.handleGetDocument).Methods(http.MethodGet)
	s.router.HandleFunc(DocumentsPath, s.handleUpdateDocument).Methods(http.MethodPut)
	s.router.HandleFunc(DocumentsPath+"/{id}", s.handleDeleteDocument).Methods(http.MethodDelete)
	s.router.HandleFunc(DocumentsPath, s.handleDocuments).Methods(http.MethodGet)
	s.router.HandleFunc(RootPath, s.handleOptions).Methods(http.MethodOptions)
}

// handleDeleteDocument позволяет удалить документ по его ID
func (s *Service) handleDeleteDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = s.engine.Delete(id)

	if err != nil {
		fmt.Println(err)

		return
	}
}

// handleUpdateDocument позволяет обновить документ (валидация отсутствует)
func (s *Service) handleUpdateDocument(w http.ResponseWriter, r *http.Request) {
	doc := model.Document{}
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.engine.Update(doc)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

// handleGetDocument позволяет получить из хранилища документ по его ID
func (s *Service) handleGetDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	doc, ok := s.engine.Get(id)
	if !ok {
		err = fmt.Errorf("Document with id `%d` is not found", id)
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	encoded, err := json.Marshal(doc)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(encoded)
	if err != nil {
		fmt.Println(err)
	}
}

// handleCreateDocument обрабатывает запрос на создание документа (валидация отсутствует)
func (s *Service) handleCreateDocument(w http.ResponseWriter, r *http.Request) {
	doc := model.Document{}
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	docWithId := s.engine.Create(doc)
	encoded, err := json.Marshal(docWithId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(encoded)
	if err != nil {
		fmt.Println(err)
	}
}

// handleDocuments обрабатывает запрос на получение всех документов из хранилища
func (s *Service) handleDocuments(w http.ResponseWriter, r *http.Request) {
	docs, err := json.Marshal(s.engine.All())
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(docs)
	if err != nil {
		fmt.Println(err)
	}
}

// handleSearch позволяет найти документы по содержимому
func (s *Service) handleSearch(w http.ResponseWriter, r *http.Request) {
	found := []model.Document{}
	token := r.URL.Query().Get("token")

	if token != "" {
		found = s.engine.Search(token)
	}

	docs, err := json.Marshal(found)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(docs)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Service) handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "OPTIONS, GET, POST, PUT, DELETE")
}
