package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/crawler/http"
	"pkg/engine"
	"pkg/engine/btree"
	"pkg/index"
	"pkg/index/word"
	"pkg/storage/file"
	"strings"
)

// Service представляет собой сервер интернет-поисковика
type Service struct {
	scanner crawler.Scanner
	index   index.Service
	engine  engine.Service
}

func main() {
	service, err := new()
	if err != nil {
		fmt.Println(err)
		return
	}

	go service.scan()
	service.readline()
}

func new() (*Service, error) {
	storage, err := file.NewStorage()
	if err != nil {
		return &Service{}, err
	}

	s := Service{
		scanner: http.NewCrawler("https://habr.com", 2),
		index:   word.NewService(storage),
		engine:  btree.NewService(storage),
	}

	return &s, nil
}

func (s *Service) scan() {
	docs, err := s.scanner.Scan()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.index.Fill(&docs)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Service) readline() {
	for {
		fmt.Println("Enter search word (leave empty to exit):")
		reader := bufio.NewReader(os.Stdin)
		word, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		word = strings.TrimSuffix(word, "\r\n")
		word = strings.TrimSuffix(word, "\n")
		if word == "" {
			break
		}

		found, err := s.engine.Search(word)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Results for '%s':\n", word)
		for _, v := range found {
			fmt.Println(v)
		}
	}

	fmt.Println("Bye!")
}
