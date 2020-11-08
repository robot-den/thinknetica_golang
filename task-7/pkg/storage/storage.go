// Package storage содержит интерфейс хранилища данных
package storage

import (
	"pkg/model"
)

// Service определяет контракт хранилища данных
type Service interface {
	Read() ([]model.Document, model.InvertedIndex, error)
	Write([]model.Document, model.InvertedIndex) error
	IsUpdated() bool
}
