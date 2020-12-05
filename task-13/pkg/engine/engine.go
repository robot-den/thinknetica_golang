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
func (s *Service) Search(query string) []*model.Document {
	ids := s.index.Search(query)
	if len(ids) == 0 {
		return []*model.Document{}
	}

	return s.storage.Read(ids)
}
