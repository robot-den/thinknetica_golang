// Package index осуществляет индексирование, хранение и поиск для результатов сканирования
package index

import (
	"fmt"
	"strings"
)

// Index предоставляет методы для индексирования, а так же временно(!) выполняет функции базы данных (хранение и поиск)
type Index struct {
	storage       BTree
	invertedIndex map[string][]int
	IdProvider    int
}

// Record представляет собой тип, хранящий данные по отдельной странице
type Record struct {
	Id    int
	Url   string
	Title string
}

// BTree представляет собой бинарное дерево, предоставляющее методы по добавлению и поиску элементов
type BTree struct {
	root *node
}

// node представляет собой узел бинарного дерева. Этот тип хранит значение и ссылки на другие узлы
type node struct {
	left  *node
	right *node
	value *Record
}

// New создает новый экземпляр типа Index
func New() *Index {
	ind := Index{
		storage:       BTree{},
		invertedIndex: map[string][]int{},
		IdProvider:    1,
	}
	return &ind
}

// Search выполняет поиск по слову в хранилище с использованием инвертированного индекса
func (i *Index) Search(word string) []string {
	found := []string{}
	ids, ok := i.invertedIndex[word]
	if !ok {
		return found
	}

	for _, id := range ids {
		if rec, ok := i.storage.Search(id); ok {
			found = append(found, fmt.Sprintf("%s - %s", rec.Url, rec.Title))
		}
	}

	return found
}

// Fill заполняет хранилище и инвертированный индекс
func (i *Index) Fill(data *map[string]string) {
	for link, title := range *data {
		rec := Record{
			Id:    i.IdProvider,
			Url:   link,
			Title: title,
		}
		i.storage.Add(&rec)

		lexemes := map[string]bool{} // map чтобы избежать дублей
		parseLexemes(&lexemes, rec.Url, "/", "#:")
		parseLexemes(&lexemes, rec.Title, " ", "()/,-")

		for lex := range lexemes {
			i.invertedIndex[lex] = append(i.invertedIndex[lex], rec.Id)
		}

		i.IdProvider++
	}
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

// Add позволяет добавить элемент в бинарное дерево
func (bt *BTree) Add(r *Record) {
	newNode := &node{
		value: r,
	}

	if bt.root == nil {
		bt.root = newNode
		return
	}

	parentNode := bt.root
	for {
		if parentNode.value.Id > r.Id {
			if parentNode.left == nil {
				parentNode.left = newNode
				break
			}
			parentNode = parentNode.left
			continue
		}

		if parentNode.right == nil {
			parentNode.right = newNode
			break
		}
		parentNode = parentNode.right
	}
}

// Search осуществляет поиск в бинарном дереве, второе возвращаемое значение равно false если запись не найдена
func (bt *BTree) Search(id int) (*Record, bool) {
	currentNode := bt.root

	for {
		if currentNode == nil {
			return &Record{}, false
		}

		if currentNode.value.Id == id {
			return currentNode.value, true
		}

		if currentNode.value.Id > id {
			currentNode = currentNode.left
			continue
		}

		currentNode = currentNode.right
	}
}

// String позволяет получить простое строковое представление бинарного дерева
func (bt *BTree) String() string {
	elems := []int{}
	bt.root.collect(&elems)
	return fmt.Sprint(elems)
}

// collect выполняет рекурсивный обход дерева и собирает Id элементов в массив
func (n *node) collect(s *[]int) {
	if n == nil {
		return
	}

	*s = append(*s, n.value.Id)

	n.left.collect(s)
	n.right.collect(s)
}
