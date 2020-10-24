// Package index осуществляет индексирование, хранение и поиск для результатов сканирования
package index

import (
	"fmt"
	"sort"
	"strings"
)

// Scanner предоставляет возможность получить словарь
type Scanner interface {
	Scan() (map[string]string, error)
}

// Index предоставляет методы для индексирования, а так же временно(!) выполняет функции базы данных (хранение и поиск)
type Index struct {
	scanner       Scanner
	storage       []Record
	invertedIndex map[string][]int
}

// Record представляет собой тип, хранящий данные по отдельной странице
type Record struct {
	Id    int
	Url   string
	Title string
}

// New создает новый экземпляр типа Index
func New(s Scanner) *Index {
	ind := Index{
		scanner:       s,
		storage:       []Record{},
		invertedIndex: map[string][]int{},
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

// Fill получает данные от scanner и заполняет хранилище и инвертированный индекс
func (i *Index) Fill() error {
	data, err := i.scanner.Scan()
	if err != nil {
		return err
	}

	i.fillStorage(&data)
	i.fillInvertedIndex()

	return nil
}

// fillStorage заполняет хранилище элементами Record и выполняет сортировку storage по record.Id
func (i *Index) fillStorage(data *map[string]string) {
	id := 1

	for link, title := range *data {
		rec := Record{
			Id:    id,
			Url:   link,
			Title: title,
		}
		id++
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
