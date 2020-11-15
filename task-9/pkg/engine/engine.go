// Package engine содержит интерфейс поискового движка
package engine

// Service определяет контракт поискового движка
type Service interface {
	Search(word string) ([]string, error)
}
