package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/engine"
	"strings"
)

func main() {
	crw := crawler.New("https://habr.com", 1)

	for {
		fmt.Println("Enter search phrase (leave empty to exit):")
		reader := bufio.NewReader(os.Stdin)
		phrase, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		phrase = strings.TrimSuffix(phrase, "\r\n")
		phrase = strings.TrimSuffix(phrase, "\n")
		if phrase == "" {
			break
		}

		found, err := engine.Search(crw, phrase)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Results for '%s':\n", phrase)
		for _, v := range found {
			fmt.Println(v)
		}
	}

	fmt.Println("Bye!")
}
