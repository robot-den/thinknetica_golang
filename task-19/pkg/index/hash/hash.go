// Package hash осуществляет индексирование и поиск документов
package hash

import (
	"strings"
	"task-19/pkg/model"
)

// Service хранит внутренний индекс документов и предоставляет операции поиска и обновления.
type Service struct {
	index map[string][]int
}

// NewService создает новый экземпляр типа Service
func NewService() *Service {
	s := Service{
		index: map[string][]int{},
	}
	return &s
}

// Search осуществляет поиск IDs документов по фразе в индексе
func (s *Service) Search(query string) []int {
	return s.index[query]
}

// Update разбирает записи model.Document на лексемы и обновляет индекс
func (s *Service) Update(docs []model.Document) {
	for _, doc := range docs {
		lexemes := map[string]bool{} // map чтобы избежать дублей
		parseLexemes(&lexemes, doc.URL, "/", "#:")
		parseLexemes(&lexemes, doc.Title, " ", "()/,-")

		for lex := range lexemes {
			s.index[lex] = append(s.index[lex], doc.ID)
		}
	}
}

// ReadAll возвращает все содержимое индекса
func (s *Service) ReadAll() map[string][]int {
	return s.index
}

// parseLexemes разбирает текст на лексемы
func parseLexemes(lexemes *map[string]bool, url, sep, tr string) {
	for _, word := range strings.Split(url, sep) {
		if tr != "" {
			word = strings.Trim(word, tr)
		}
		// Не храним данные по "словам" короче 2х символов ¯\_(ツ)_/¯
		if len([]rune(word)) > 1 {
			(*lexemes)[word] = true
		}
	}
}
