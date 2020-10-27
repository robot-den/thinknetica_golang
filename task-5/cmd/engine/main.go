package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/index"
	"strings"
)

func main() {
	// В production сканирование сайтов и индексирование результатов выполнялось бы отдельно от сервиса поиска
	fmt.Println("Scanning sites and indexing results...")
	crw := crawler.New("https://habr.com", 2)
	webData, err := crw.Scan()
	if err != nil {
		fmt.Println(err)
		return
	}
	ind := index.New()
	ind.Fill(&webData)

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

		found := ind.Search(word)
		fmt.Printf("Results for '%s':\n", word)
		for _, v := range found {
			fmt.Println(v)
		}
	}

	fmt.Println("Bye!")
}
