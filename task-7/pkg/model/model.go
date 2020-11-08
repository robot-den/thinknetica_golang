// Package model предоставляет структуры данных которые используются в нескольких пакетах приложения
package model

// Document представляет собой проиндексированную веб-страницу
type Document struct {
	ID    int
	URL   string
	Title string
}

// InvertedIndex представляет собой оборатный индекс в котором ключ это слово, а значение это массив Document.ID
type InvertedIndex map[string][]int
