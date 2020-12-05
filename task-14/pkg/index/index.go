// Package index содержит интерфейс индексатора документов
package index

import "pkg/model"

// Service определяет контракт индексатора документов
type Service interface {
	Update(data []*model.Document)
	Search(string) []int
}
