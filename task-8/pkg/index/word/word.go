// Package word осуществляет индексирование, хранение и поиск для результатов сканирования
package word

import (
	"pkg/model"
	"strings"
)

// Service создает записи типа model.Document, формирует обратный индекс и сохраняет результаты в указанное хранилище
type Service struct {
	storage    writer
	IDProvider int
}

// writer представляет собой интерфейс (контракт) которому должно соответствовать хранилище
type writer interface {
	Write([]model.Document, model.InvertedIndex) error
}

// NewService создает новый экземпляр типа Service
func NewService(storage writer) *Service {
	ind := Service{
		storage:    storage,
		IDProvider: 1,
	}
	return &ind
}

// Fill помещает в хранилище массив записей типа model.Document и обратный индекс
func (i *Service) Fill(data *map[string]string) error {
	docs := []model.Document{}
	invertedIndex := model.InvertedIndex{}

	for link, title := range *data {
		doc := model.Document{
			ID:    i.IDProvider,
			URL:   link,
			Title: title,
		}
		docs = append(docs, doc)

		lexemes := map[string]bool{} // map чтобы избежать дублей
		parseLexemes(&lexemes, doc.URL, "/", "#:")
		parseLexemes(&lexemes, doc.Title, " ", "()/,-")

		for lex := range lexemes {
			invertedIndex[lex] = append(invertedIndex[lex], doc.ID)
		}

		i.IDProvider++
	}

	err := i.storage.Write(docs, invertedIndex)
	if err != nil {
		return err
	}

	return nil
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
