// package rpcsrv позволяет клиентам работать с документами при помощи RPC
package rpcsrv

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"task-19/pkg/engine"
	"task-19/pkg/model"
)

// RpcSrv представляет собой RPC сервис
type RpcSrv struct {
	server *Server
	port   string
}

// New создает новый объект RPC сервиса
func New(eng *engine.Service, port string) *RpcSrv {
	r := RpcSrv{
		server: &Server{engine: eng},
		port:   port,
	}
	return &r
}

// Serve запускает сервер, обслуживающий RPC запросы
func (r *RpcSrv) Serve() {
	err := rpc.Register(r.server)
	if err != nil {
		fmt.Println(err)
		return
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", r.port)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = http.Serve(listener, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// Server представляет собой экспортируемый тип данных, реализующий методы, доступные для удаленного вызова
type Server struct {
	engine *engine.Service
}

// Search выполняет поиск документов по поисковому запросу
func (s *Server) Search(token string, documents *[]model.Document) error {
	*documents = s.engine.Search(token)
	return nil
}
