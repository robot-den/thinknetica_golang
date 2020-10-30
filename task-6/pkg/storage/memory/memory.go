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
	records        []model.Record
	index          model.InvertedIndex
}

// NewStorage создает новое хранилище
func NewStorage() (*Storage, error) {
	s := Storage{
		checksum:       [16]byte{},
		clientChecksum: [16]byte{},
		records:        []model.Record{},
		index:          model.InvertedIndex{},
	}
	return &s, nil
}

// WriteRecords сохраняет массив записей model.Record и индекс в поля структуры Storage
func (s *Storage) Write(records []model.Record, index model.InvertedIndex) error {
	s.records = records
	s.index = index
	s.checksum = md5.Sum([]byte(fmt.Sprintf("%v", records)))

	return nil
}

// ReadRecords читает массив записей model.Record из поля структуры Storage
func (s *Storage) Read() ([]model.Record, model.InvertedIndex, error) {
	// Обновляем каждый раз когда клиент читает данные
	s.clientChecksum = s.checksum

	return s.records, s.index, nil
}

// IsUpdated позволяет клиенту проверить обновились ли данные после последнего чтения
func (s *Storage) IsUpdated() bool {
	return s.clientChecksum != s.checksum
}
