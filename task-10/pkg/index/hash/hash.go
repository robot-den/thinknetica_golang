package hash

import (
	"pkg/crawler"
	"strings"
)

// Index - индекс на основе хэш-таблицы.
type Index struct {
	data map[string][]int
}

// New - конструктор.
func New() *Index {
	var ind Index
	ind.data = make(map[string][]int)
	return &ind
}

// Add добавляет данные из переданных документов в индекс.
//
// Сначала происходит выделение лексем как ключей словаря из данных документа.
// Потом проверяется наличие номера документа в занчении словаря для лексемы.
// Если номер документа не найден, то он добавляется в значение словаря.
// Возвращает число новых ключей, созданных в процессе добавления документов.
func (index *Index) Add(docs []crawler.Document) int {
	var newKeysCount int
	for _, doc := range docs {
		for _, token := range tokens(doc.Title) {
			if !exists(index.data[token], doc.ID) {
				ids, ok := index.data[token]
				if !ok {
					newKeysCount++
				}
				index.data[token] = append(ids, doc.ID)
			}
		}
	}
	return newKeysCount
}

// Search возвращает номера документов, где встречается данная лексема.
func (index *Index) Search(token string) []int {
	return index.data[strings.ToLower(token)]
}

// Разделение строки на лексемы.
func tokens(s string) []string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

// Проверка наличия элемента в массиве.
func exists(ids []int, item int) bool {
	for _, id := range ids {
		if id == item {
			return true
		}
	}
	return false
}
