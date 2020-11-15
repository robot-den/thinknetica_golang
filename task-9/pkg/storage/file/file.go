// Package file предоставляет возможность сохранить данные в файл
package file

import (
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"os"
	"pkg/model"
)

// DocumentsFile путь к файлу который содержит сериализованные записи типа Document
const DocumentsFile = "./documents.json"

// IndexFile путь к файлу который содержит сериализованный обратный индекс
const IndexFile = "./index.json"

// Storage - тип реализующий методы чтения/записи данных
type Storage struct {
	checksum       [16]byte
	clientChecksum [16]byte
}

// NewStorage создает новое файловое хранилище
func NewStorage() (*Storage, error) {
	s := Storage{
		clientChecksum: [16]byte{},
	}
	c, err := checksum()
	if err != nil {
		return &s, err
	}
	s.checksum = c
	return &s, nil
}

// Write сериализует в JSON и записывает в файл массив model.Document и обратный индекс
func (s *Storage) Write(documents []model.Document, index model.InvertedIndex) error {
	err := s.writeDocuments(documents)
	if err != nil {
		return err
	}
	err = s.writeInvertedIndex(index)
	if err != nil {
		return err
	}

	return nil
}

// Read читает из файла массив model.Document и обратный индекс и десериализует их
func (s *Storage) Read() ([]model.Document, model.InvertedIndex, error) {
	documents, err := s.readDocuments()
	if err != nil {
		return documents, model.InvertedIndex{}, err
	}
	index, err := s.readInvertedIndex()
	if err != nil {
		return documents, index, err
	}

	return documents, index, nil
}

// IsUpdated позволяет клиенту проверить обновились ли данные после последнего чтения
func (s *Storage) IsUpdated() bool {
	return s.clientChecksum != s.checksum
}

// WriteDocuments сериализует в JSON и записывает в файл массив model.Document
func (s *Storage) writeDocuments(documents []model.Document) error {
	encodedDocuments, err := json.Marshal(documents)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(DocumentsFile, encodedDocuments, 0666)
	if err != nil {
		return err
	}
	s.checksum = md5.Sum(encodedDocuments)

	return nil
}

// WriteInvertedIndex сериализует в JSON и записывает в файл обратный индекс
func (s *Storage) writeInvertedIndex(index model.InvertedIndex) error {
	encodedIndex, err := json.Marshal(index)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(IndexFile, encodedIndex, 0666)
	if err != nil {
		return err
	}

	return nil
}

// ReadDocuments читает из файла JSON (массив записей типа model.Document) и десериализует его
func (s *Storage) readDocuments() ([]model.Document, error) {
	documents := []model.Document{}

	encodedDocuments, err := ioutil.ReadFile(DocumentsFile)
	if err != nil {
		// Поскольку мы не разрешаем менять расположение файла с данными, его отсутствие не является ошибкой
		if os.IsNotExist(err) {
			return documents, nil
		}
		return documents, err
	}

	err = json.Unmarshal(encodedDocuments, &documents)
	if err != nil {
		return documents, err
	}
	// Обновляем каждый раз когда клиент читает данные
	s.clientChecksum = s.checksum

	return documents, nil
}

// ReadInvertedIndex читает из файла JSON (обратный индекс) и десериализует его
func (s *Storage) readInvertedIndex() (model.InvertedIndex, error) {
	index := model.InvertedIndex{}

	encodedIndex, err := ioutil.ReadFile(IndexFile)
	if err != nil {
		// Поскольку мы не разрешаем менять расположение файла с данными, его отсутствие не является ошибкой
		if os.IsNotExist(err) {
			return index, nil
		}
		return index, err
	}

	err = json.Unmarshal(encodedIndex, &index)
	if err != nil {
		return index, err
	}

	return index, nil
}

func checksum() ([16]byte, error) {
	data, err := ioutil.ReadFile(DocumentsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return [16]byte{}, nil
		}
		return [16]byte{}, err
	}

	sum := md5.Sum(data)

	return sum, nil
}
