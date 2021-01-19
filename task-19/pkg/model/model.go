// Package model предоставляет структуры данных которые используются в нескольких пакетах приложения
package model

// Document представляет собой проиндексированную веб-страницу
type Document struct {
	ID    int
	URL   string
	Title string
}
