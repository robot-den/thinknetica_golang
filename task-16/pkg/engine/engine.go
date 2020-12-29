// Package engine реализует поисковый движок
package engine

import (
	"pkg/index"
	"pkg/model"
	"pkg/storage"
)

// Service определяет контракт поискового движка
type Service struct {
	index   index.Service
	storage storage.Service
}

// NewService создает новый экземпляр типа Service
func NewService(i index.Service, s storage.Service) *Service {
	service := Service{
		index:   i,
		storage: s,
	}
	return &service
}

// Search выполняет поиск по фразе.
// Сначала получает IDs документов в индексе, а затем по ним ищет записи в хранилище.
func (s *Service) Search(query string) []model.Document {
	ids := s.index.Search(query)
	if len(ids) == 0 {
		return []model.Document{}
	}

	return s.storage.Read(ids)
}

// All позволяет получить все документы из хранилища
func (s *Service) All() []model.Document {
	return s.storage.ReadAll()
}

// BatchCreate позволяет сохранить массив документов
func (s *Service) BatchCreate(docs []model.Document) []model.Document {
	docsWithIds := s.storage.Write(docs)
	s.index.Update(docsWithIds)
	return docsWithIds
}

// Create позволяет сохранить новый документ в хранилище и индексе
func (s *Service) Create(doc model.Document) model.Document {
	docs := []model.Document{doc}
	docsWithIds := s.storage.Write(docs)
	s.index.Update(docsWithIds)
	return docsWithIds[0]
}

// Get позволяет получить документ по его ID (возвращает также маркер указывающий что документ найден)
func (s *Service) Get(id int) (model.Document, bool) {
	ids := []int{id}
	docs := s.storage.Read(ids)
	if len(docs) <= 0 {
		return model.Document{}, false
	}

	return docs[0], true
}

// Update позволяет обновить документ (целостность индекса нарушается)
func (s *Service) Update(doc model.Document) error {
	err := s.storage.Update(doc)
	return err
}

// Delete позволяет удалить документ по его ID (целостность индекса нарушается)
func (s *Service) Delete(id int) error {
	err := s.storage.Delete(id)
	return err
}
