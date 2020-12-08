// Package plugin содержит интерфейс подключаемых расширений
package plugin

// Service определяет контракт подключаемого расширения
type Service interface {
	Run()
}
