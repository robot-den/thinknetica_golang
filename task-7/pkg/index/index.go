// Package index содержит интерфейс индексатора документов
package index

// Service определяет контракт индексатора документов
type Service interface {
	Fill(data *map[string]string) error
}
