// Package file предоставляет возможность сохранить данные в файл
package file

import (
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"os"
	"pkg/model"
)

// RecordsFile путь к файлу который содержит сериализованные записи типа Record
const RecordsFile = "./records.json"

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

// Write сериализует в JSON и записывает в файл массив model.Record и обратный индекс
func (s *Storage) Write(records []model.Record, index model.InvertedIndex) error {
	err := s.writeRecords(records)
	if err != nil {
		return err
	}
	err = s.writeInvertedIndex(index)
	if err != nil {
		return err
	}

	return nil
}

// Read читает из файла массив model.Record и обратный индекс и десериализует их
func (s *Storage) Read() ([]model.Record, model.InvertedIndex, error) {
	records, err := s.readRecords()
	if err != nil {
		return records, model.InvertedIndex{}, err
	}
	index, err := s.readInvertedIndex()
	if err != nil {
		return records, index, err
	}

	return records, index, nil
}

// WriteRecords сериализует в JSON и записывает в файл массив model.Record
func (s *Storage) writeRecords(records []model.Record) error {
	encodedRecords, err := json.Marshal(records)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(RecordsFile, encodedRecords, 0666)
	if err != nil {
		return err
	}
	s.checksum = md5.Sum(encodedRecords)

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

// ReadRecords читает из файла JSON (массив записей типа model.Record) и десериализует его
func (s *Storage) readRecords() ([]model.Record, error) {
	records := []model.Record{}

	encodedRecords, err := ioutil.ReadFile(RecordsFile)
	if err != nil {
		// Поскольку мы не разрешаем менять расположение файла с данными, его отсутствие не является ошибкой
		if os.IsNotExist(err) {
			return records, nil
		}
		return records, err
	}

	err = json.Unmarshal(encodedRecords, &records)
	if err != nil {
		return records, err
	}
	// Обновляем каждый раз когда клиент читает данные
	s.clientChecksum = s.checksum

	return records, nil
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

// IsUpdated позволяет клиенту проверить обновились ли данные после последнего чтения
func (s *Storage) IsUpdated() bool {
	return s.clientChecksum != s.checksum
}

func checksum() ([16]byte, error) {
	data, err := ioutil.ReadFile(RecordsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return [16]byte{}, nil
		}
		return [16]byte{}, err
	}

	sum := md5.Sum(data)

	return sum, nil
}
