// Package memory предоставляет возможность сохранить документы в памяти и получить их по ID
package memory

import (
	"fmt"
	"pkg/model"
	"sort"
	"sync"
)

// Storage - тип реализующий методы чтения/записи документов
type Storage struct {
	documents  []model.Document
	IDProvider int
	mu         *sync.Mutex
}

// NewStorage создает новое хранилище
func NewStorage() *Storage {
	s := Storage{
		documents:  []model.Document{},
		IDProvider: 1,
		mu:         &sync.Mutex{},
	}
	return &s
}

// Read возвращает массив записей model.Document по их IDs
func (s *Storage) Read(ids []int) []model.Document {
	s.mu.Lock()
	defer s.mu.Unlock()

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
	s.mu.Lock()
	defer s.mu.Unlock()

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

// Update позволяет обновить указанный документ
func (s *Storage) Update(doc model.Document) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	notFound := fmt.Errorf("document with id `%d` is not found", doc.ID)

	docsCount := len(s.documents)
	index := sort.Search(docsCount, func(ind int) bool {
		return (s.documents[ind]).ID >= doc.ID
	})

	if index < docsCount {
		foundDoc := s.documents[index]
		if foundDoc.ID == doc.ID {
			notFound = nil
			s.documents[index] = doc
		}
	}

	return notFound
}

// Delete позволяет удалить указанный документ по его ID
func (s *Storage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	notFound := fmt.Errorf("document with id `%d` is not found", id)

	docsCount := len(s.documents)
	index := sort.Search(docsCount, func(ind int) bool {
		return (s.documents[ind]).ID >= id
	})

	if index < docsCount {
		foundDoc := s.documents[index]
		if foundDoc.ID == id {
			notFound = nil
			s.documents = append(s.documents[:index], s.documents[index+1:]...)
		}
	}

	return notFound
}

// ReadAll возвращает все документы содержащиеся в хранилище
func (s *Storage) ReadAll() []model.Document {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.documents
}
