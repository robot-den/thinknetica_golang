package index

// Обратный индекс отсканированных документов.

import "pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document) int
	Search(string) []int
}
