// Package memory предоставляет возможность сохранить данные в памяти
package memory

import (
	"crypto/md5"
	"fmt"
	"pkg/model"
)

// Storage - тип реализующий методы чтения/записи данных
type Storage struct {
	checksum       [16]byte
	clientChecksum [16]byte
	documents      []model.Document
	index          model.InvertedIndex
}

// NewStorage создает новое хранилище
func NewStorage() (*Storage, error) {
	s := Storage{
		checksum:       [16]byte{},
		clientChecksum: [16]byte{},
		documents:      []model.Document{},
		index:          model.InvertedIndex{},
	}
	return &s, nil
}

// WriteDocuments сохраняет массив записей model.Document и индекс в поля структуры Storage
func (s *Storage) Write(documents []model.Document, index model.InvertedIndex) error {
	s.documents = documents
	s.index = index
	s.checksum = md5.Sum([]byte(fmt.Sprintf("%v", documents)))

	return nil
}

// ReadDocuments читает массив записей model.Document из поля структуры Storage
func (s *Storage) Read() ([]model.Document, model.InvertedIndex, error) {
	// Обновляем каждый раз когда клиент читает данные
	s.clientChecksum = s.checksum

	return s.documents, s.index, nil
}

// IsUpdated позволяет клиенту проверить обновились ли данные после последнего чтения
func (s *Storage) IsUpdated() bool {
	return s.clientChecksum != s.checksum
}
