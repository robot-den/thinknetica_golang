package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"strings"
)

// Service представляет собой CLI клиент
type Service struct {
	reader *bufio.Reader
}

func new() *Service {
	s := Service{
		reader: bufio.NewReader(os.Stdin),
	}
	return &s
}

func main() {
	service := new()
	go service.subscribe()
	service.interact()
}

// subscribe позволяет подписаться на сообщения
func (s *Service) subscribe() {
	conn, r, err := websocket.DefaultDialer.Dial("ws://localhost:9000/messages", nil)
	if err != nil {
		fmt.Println(err, r.StatusCode)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			fmt.Println(err)
			return
		}
		fmt.Println(string(message))
	}
}

// interact позволяет отправлять сообщения другим участникам чата (после аутентификации)
func (s *Service) interact() {
	fmt.Println("Enter password to start chatting:")

	password, err := s.readline()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Start chatting:")
	for {
		message, err := s.readline()
		if err != nil {
			fmt.Println(err)
			return
		}

		if message == "" {
			break
		}

		err = s.sendMessage(password, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Bye!")
}

// sendMessage выполняет отправку сообщения на сервер чата
func (s *Service) sendMessage(password, message string) error {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:9000/send", nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte(password))
	if err != nil {
		return err
	}

	_, resp, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	if string(resp) != "OK" {
		return fmt.Errorf("you aren't authenticated")
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// readline позволяет построчно принимать сообщения от пользователя
func (s *Service) readline() (string, error) {
	str, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	str = strings.TrimSuffix(str, "\r\n")
	str = strings.TrimSuffix(str, "\n")

	return str, nil
}
