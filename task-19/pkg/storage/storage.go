// Package storage содержит интерфейс хранилища данных
package storage

import (
	"task-19/pkg/model"
)

// Service определяет контракт хранилища данных
type Service interface {
	Read([]int) []model.Document
	Write([]model.Document) []model.Document
	Update(model.Document) error
	Delete(int) error
	ReadAll() []model.Document
}
