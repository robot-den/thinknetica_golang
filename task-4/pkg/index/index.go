// Package index осуществляет индексирование, хранение и поиск для результатов сканирования
package index

import (
	"fmt"
	"sort"
	"strings"
)

// Index предоставляет методы для индексирования, а так же временно(!) выполняет функции базы данных (хранение и поиск)
type Index struct {
	storage       []Record
	invertedIndex map[string][]int
	IdProvider    int
}

// Record представляет собой тип, хранящий данные по отдельной странице
type Record struct {
	Id    int
	Url   string
	Title string
}

// New создает новый экземпляр типа Index
func New() *Index {
	ind := Index{
		storage:       []Record{},
		invertedIndex: map[string][]int{},
		IdProvider:    1,
	}
	return &ind
}

// Search выполняет поиск по слову в хранилище с использованием инвертированного индекса
// используется sort.Search(), который требует дополнительной проверки что коллекция содержит искомый элемент
func (i *Index) Search(word string) []string {
	found := []string{}
	storageLength := len(i.storage)
	ids, ok := i.invertedIndex[word]
	if !ok {
		return found
	}

	for _, id := range ids {
		index := sort.Search(storageLength, func(ind int) bool {
			return (i.storage[ind]).Id >= id
		})

		if index < storageLength {
			rec := i.storage[index]
			if rec.Id == id {
				found = append(found, fmt.Sprintf("%s - %s", rec.Url, rec.Title))
			}
		}
	}

	return found
}

// Fill заполняет хранилище и инвертированный индекс
func (i *Index) Fill(data *map[string]string) {
	i.fillStorage(data)
	i.fillInvertedIndex()
}

// fillStorage заполняет хранилище элементами Record и выполняет сортировку storage по record.Id
func (i *Index) fillStorage(data *map[string]string) {
	for link, title := range *data {
		rec := Record{
			Id:    i.IdProvider,
			Url:   link,
			Title: title,
		}
		i.IdProvider++
		i.storage = append(i.storage, rec)
	}

	// Хотя данные уже отсортированы, все равно сортируем чтобы поиграть с сортировкой
	sort.Slice(i.storage, func(a, b int) bool { return i.storage[a].Id < i.storage[b].Id })
}

// fillInvertedIndex разбирает record.Url и record.Title на примитивные лексемы и заполняет индекс
// Индекс это ассоциативный массив, где ключ это лексема, а значение это массив состоящий из record.Id
func (i *Index) fillInvertedIndex() {
	for _, record := range i.storage {
		lexemes := map[string]bool{} // map чтобы избежать дублей
		parseLexemes(&lexemes, record.Url, "/", "#:")
		parseLexemes(&lexemes, record.Title, " ", "()/,-")

		for lex := range lexemes {
			i.invertedIndex[lex] = append(i.invertedIndex[lex], record.Id)
		}
	}
}

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
