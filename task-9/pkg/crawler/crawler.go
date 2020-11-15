// Package crawler содержит интерфейс поискового робота
package crawler

// Scanner определяет контракт поискового робота
type Scanner interface {
	Scan() (data map[string]string, err error)
}
