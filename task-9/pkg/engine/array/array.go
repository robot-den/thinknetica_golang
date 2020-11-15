// Package array читает данные из хранилища и предоставляет к ним доступ с использованием отсортированного массива
package array

import (
	"fmt"
	"pkg/model"
	"sort"
)

// Service предоставляет методы для поиска данных в указанном хранилище
type Service struct {
	storage       storage
	arr           []*model.Document
	invertedIndex model.InvertedIndex
}

// storage представляет собой интерфейс (контракт) которому должно соответствовать хранилище
type storage interface {
	Read() ([]model.Document, model.InvertedIndex, error)
	IsUpdated() bool
}

// NewService создает новый экземпляр типа Service
func NewService(s storage) *Service {
	eng := Service{
		storage: s,
		arr:     []*model.Document{},
	}
	return &eng
}

// Search осуществляет поиск по слову
func (s *Service) Search(word string) ([]string, error) {
	if s.storage.IsUpdated() {
		err := s.update()
		if err != nil {
			return []string{}, err
		}
	}

	found := []string{}
	ids, ok := s.invertedIndex[word]
	if !ok {
		return found, nil
	}

	// TODO: replace search in btree with search in sorted array; adjust specs, write 2 banchmarks - for btree and for array
	storageLength := len(s.arr)
	for _, id := range ids {
		index := sort.Search(storageLength, func(ind int) bool {
			return (s.arr[ind]).ID >= id
		})

		if index < storageLength {
			rec := s.arr[index]
			if rec.ID == id {
				found = append(found, fmt.Sprintf("%s - %s", rec.URL, rec.Title))
			}
		}
	}

	return found, nil
}

// update читает данные из хранилища и строит в памяти необходимые структуры для быстрого доступа к ним
func (s *Service) update() error {
	documents, index, err := s.storage.Read()
	if err != nil {
		return err
	}

	arr := []*model.Document{}
	for _, doc := range documents {
		document := doc // приходится копировать здесь, иначе все записи ссылаются на одно и то же значение
		arr = append(arr, &document)
	}
	sort.Slice(arr, func(a, b int) bool { return arr[a].ID < arr[b].ID })
	s.arr = arr
	s.invertedIndex = index

	return nil
}
