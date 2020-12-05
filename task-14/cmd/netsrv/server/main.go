package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"pkg/crawler/webscnr"
	"pkg/engine"
	"pkg/index"
	"pkg/index/hash"
	"pkg/storage"
	"pkg/storage/memory"
	"strings"
	"time"
)

// Service представляет собой сервер интернет-поисковика
type Service struct {
	storage storage.Service
	index   index.Service
	engine  *engine.Service
}

func main() {
	service := new()

	go service.scan()
	service.serveRequests()
}

func new() *Service {
	str := memory.NewStorage()
	ind := hash.NewService()
	s := Service{
		storage: str,
		index:   ind,
		engine:  engine.NewService(ind, str),
	}

	return &s
}

// scan сканирует сайты, сохраняет результаты в хранилище и обновляет индекс.
func (s *Service) scan() {
	scnr := &webscnr.WebScnr{}
	docs, err := scnr.Scan("https://redis.io/", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	docsWithIds := s.storage.Write(docs)
	s.index.Update(docsWithIds)
}

// serveRequests запускает службу для обслуживания запросов
func (s *Service) serveRequests() {
	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server is waiting for connections")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go s.handleConn(conn)
	}
}

// handleConn обрабатывает запросы клиента: ищет документы по запросу и возвращает массив в формате JSON
func (s *Service) handleConn(conn net.Conn) {
	defer conn.Close()

	connReader := bufio.NewReader(conn)

	fmt.Println("Handle connection:", conn.RemoteAddr(), time.Now())
	for {
		err := conn.SetDeadline(time.Now().Add(time.Second * 10))
		if err != nil {
			fmt.Println(err)
			return
		}

		query, err := connReader.ReadString('\n')
		if err != nil {
			// do not notify when client disconnects
			if err != io.EOF {
				fmt.Println(err)
			}
			return
		}
		query = strings.TrimSuffix(query, "\r\n")
		query = strings.TrimSuffix(query, "\n")
		if query == "" {
			break
		}

		found := s.engine.Search(query)
		resp, err := json.Marshal(found)
		if err != nil {
			fmt.Println(err)
			return
		}
		// append new line character to let client know where message is ended
		resp = append(resp, '\n')

		_, err = conn.Write(resp)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
