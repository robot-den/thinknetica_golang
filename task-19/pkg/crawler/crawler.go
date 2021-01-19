// Package crawler содержит интерфейс поискового робота
package crawler

import "task-19/pkg/model"

// Scanner определяет контракт поискового робота
type Scanner interface {
	Scan(string, int) ([]model.Document, error)
}
