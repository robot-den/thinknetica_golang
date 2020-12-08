package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Connect to remote service...")
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	cliReader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	fmt.Println("Enter search query (leave empty to exit):")
	for {
		query, err := cliReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if query == "\r\n" || query == "\n" {
			break
		}

		_, err = conn.Write([]byte(query))
		if err != nil {
			fmt.Println(err)
			return
		}

		msg, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(msg)
	}

	fmt.Println("Bye!")
}
