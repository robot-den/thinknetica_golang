// package netsrv реализует подключаемый плагин, который обслуживает поисковые запросы
package netsrv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"pkg/model"
	"strings"
	"time"
)

// storage представляет собой контракт хранилища, в котором плагин netsrv будет осуществлять поиск документов
type storage interface {
	Search(string) []model.Document
}
// NetSrv представляет собой плагин
type NetSrv struct {
	storage storage
	network string
	address string
}
// New позволяет создать новый объект плагина с заданными настройками
func New(serv storage, network, address string) *NetSrv {
	s := NetSrv{
		storage: serv,
		network: network,
		address: address,
	}

	return &s
}

// Run позволяет плагину соответствовать интерфейсу plugin.Service
func (s *NetSrv) Run() {
	s.serveRequests()
}

// serveRequests запускает службу для обслуживания запросов
func (s *NetSrv) serveRequests() {
	listener, err := net.Listen(s.network, s.address)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Server is waiting for connections")
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
func (s *NetSrv) handleConn(conn net.Conn) {
	defer conn.Close()

	connReader := bufio.NewReader(conn)

	//fmt.Println("Handle connection:", conn.RemoteAddr(), time.Now())
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

		found := s.storage.Search(query)
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
