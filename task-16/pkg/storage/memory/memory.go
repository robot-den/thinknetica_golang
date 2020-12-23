// Package memory предоставляет возможность сохранить документы в памяти и получить их по ID
package memory

import (
	"pkg/model"
	"sort"
)

// Storage - тип реализующий методы чтения/записи документов
type Storage struct {
	documents  []model.Document
	IDProvider int
}

// NewStorage создает новое хранилище
func NewStorage() *Storage {
	s := Storage{
		documents:  []model.Document{},
		IDProvider: 1,
	}
	return &s
}

// Read возвращает массив записей model.Document по их IDs
func (s *Storage) Read(ids []int) []model.Document {
	var found []model.Document
	docsCount := len(s.documents)

	for _, id := range ids {
		index := sort.Search(docsCount, func(ind int) bool {
			return (s.documents[ind]).ID >= id
		})

		if index < docsCount {
			doc := s.documents[index]
			if doc.ID == id {
				found = append(found, doc)
			}
		}
	}

	return found
}

// Write сохраняет записи model.Document с присвоением им ID
func (s *Storage) Write(documents []model.Document) []model.Document {
	var docsWithIDs []model.Document
	for _, doc := range documents {
		doc.ID = s.IDProvider
		docsWithIDs = append(docsWithIDs, doc)
		s.IDProvider++
	}

	s.documents = append(s.documents, docsWithIDs...)

	// NOTE: sort is useless at the moment (all documents already sorted by ID)
	sort.Slice(s.documents, func(i, j int) bool { return s.documents[i].ID < s.documents[j].ID })

	return docsWithIDs
}

// ReadAll возвращает все документы содержащиеся в хранилище
func (s *Storage) ReadAll() []model.Document {
	return s.documents
}
