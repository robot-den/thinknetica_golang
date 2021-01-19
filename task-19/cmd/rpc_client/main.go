package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
	"task-19/pkg/model"
)

func main() {
	readline()
}

// readline читает запрос пользователя и выполняет поиск документов
func readline() {
	for {
		fmt.Println("Enter search token (leave empty to exit):")
		reader := bufio.NewReader(os.Stdin)
		token, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		token = strings.TrimSuffix(token, "\r\n")
		token = strings.TrimSuffix(token, "\n")
		if token == "" {
			break
		}

		found, err := search(token)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Results for '%s':\n", token)
		for _, v := range found {
			fmt.Println(v)
		}
	}

	fmt.Println("Bye!")
}

// search выполняет поиск документов на сервере с помощью RPC
func search(token string) ([]model.Document, error) {
	client, err := rpc.DialHTTP("tcp", ":9100")
	if err != nil {
		return []model.Document{}, err
	}
	defer client.Close()

	var documents []model.Document
	err = client.Call("Server.Search", token, &documents)
	if err != nil {
		return []model.Document{}, err
	}

	return documents, nil
}
