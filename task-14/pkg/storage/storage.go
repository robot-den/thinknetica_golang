// Package storage содержит интерфейс хранилища данных
package storage

import (
	"pkg/model"
)

// Service определяет контракт хранилища данных
type Service interface {
	Read([]int) []*model.Document
	Write([]*model.Document) []*model.Document
}
