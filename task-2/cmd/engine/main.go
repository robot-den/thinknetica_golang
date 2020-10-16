package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"strings"
)

func main() {
	url := "https://habr.com"
	depth := 2
	fmt.Printf("Scanning '%s'...\n", url)
	titles, err := crawler.Scan(url, depth)
	if err != nil {
		fmt.Println(err)
		return
	}

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

		fmt.Printf("Results for '%s':\n", phrase)
		for k, v := range titles {
			if strings.Contains(k, phrase) || strings.Contains(v, phrase) {
				fmt.Printf("%s - '%s'\n", k, v)
			}
		}
		fmt.Println()
	}

	fmt.Println("Bye!")
}
