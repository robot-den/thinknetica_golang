package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/engine"
	"pkg/index"
	"pkg/storage/file"
	"strings"
)

func main() {
	storage, err := file.NewStorage()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Running DB updating in separate process...")
	go collectUpdates(storage)

	eng := engine.New(storage)

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

		found, err := eng.Search(word)
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

func collectUpdates(storage *file.Storage) {
	crw := crawler.New("https://habr.com", 2)
	webData, err := crw.Scan()
	if err != nil {
		fmt.Println(err)
		return
	}

	ind := index.New(storage)
	err = ind.Fill(&webData)
	if err != nil {
		fmt.Println(err)
		return
	}
}
