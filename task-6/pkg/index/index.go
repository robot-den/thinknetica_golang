// Package index осуществляет индексирование, хранение и поиск для результатов сканирования
package index

import (
	"pkg/model"
	"strings"
)

// Index создает записи типа model.Record, формирует обратный индекс и сохраняет результаты в указанное хранилище
type Index struct {
	storage    writer
	IdProvider int
}

// writer представляет собой интерфейс (контракт) которому должно соответствовать хранилище
type writer interface {
	Write([]model.Record, model.InvertedIndex) error
}

// New создает новый экземпляр типа Index
func New(storage writer) *Index {
	ind := Index{
		storage:    storage,
		IdProvider: 1,
	}
	return &ind
}

// Fill помещает в хранилище массив записей типа model.Record и обратный индекс
func (i *Index) Fill(data *map[string]string) error {
	records := []model.Record{}
	invertedIndex := model.InvertedIndex{}

	for link, title := range *data {
		rec := model.Record{
			Id:    i.IdProvider,
			Url:   link,
			Title: title,
		}
		records = append(records, rec)

		lexemes := map[string]bool{} // map чтобы избежать дублей
		parseLexemes(&lexemes, rec.Url, "/", "#:")
		parseLexemes(&lexemes, rec.Title, " ", "()/,-")

		for lex := range lexemes {
			invertedIndex[lex] = append(invertedIndex[lex], rec.Id)
		}

		i.IdProvider++
	}

	err := i.storage.Write(records, invertedIndex)
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
