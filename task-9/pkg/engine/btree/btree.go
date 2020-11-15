// Package btree читает данные из хранилища и предоставляет к ним доступ с использованием бинарного дерева
package btree

import (
	"fmt"
	"pkg/btree"
	"pkg/model"
)

// Service предоставляет методы для поиска данных в указанном хранилище
type Service struct {
	storage       storage
	tree          btree.BTree
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
		tree:    btree.BTree{},
	}
	return &eng
}

// Search осуществляет поиск по слову
func (e *Service) Search(word string) ([]string, error) {
	if e.storage.IsUpdated() {
		err := e.update()
		if err != nil {
			return []string{}, err
		}
	}

	found := []string{}
	ids, ok := e.invertedIndex[word]
	if !ok {
		return found, nil
	}

	for _, id := range ids {
		if doc, ok := e.tree.Search(id); ok {
			found = append(found, fmt.Sprintf("%s - %s", doc.URL, doc.Title))
		}
	}

	return found, nil
}

// update читает данные из хранилища и строит в памяти необходимые структуры для быстрого доступа к ним
func (e *Service) update() error {
	documents, index, err := e.storage.Read()
	if err != nil {
		return err
	}

	tr := btree.BTree{}
	for _, doc := range documents {
		document := doc // приходится копировать здесь, иначе все записи ссылаются на одно и то же значение
		tr.Add(&document)
	}
	e.tree = tr
	e.invertedIndex = index

	return nil
}
