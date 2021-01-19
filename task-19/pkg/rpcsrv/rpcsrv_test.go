package rpcsrv

import (
	"net/http/httptest"
	"net/rpc"
	"reflect"
	"task-19/pkg/engine"
	"task-19/pkg/index/hash"
	"task-19/pkg/model"
	"task-19/pkg/storage/memory"
	"testing"
)

func TestRpcSrv_Serve(t *testing.T) {
	str := memory.NewStorage()
	ind := hash.NewService()
	eng := engine.NewService(ind, str)
	docs := []model.Document{
		{URL: "URL 1", Title: "Title1"},
		{URL: "URL 2", Title: "Title2"},
	}
	eng.BatchCreate(docs)
	srv := &Server{
		engine: eng,
	}

	_ = rpc.Register(srv)
	rpc.HandleHTTP()
	server := httptest.NewServer(nil)

	client, err := rpc.DialHTTP("tcp", server.Listener.Addr().String())
	if err != nil {
		t.Errorf("rpc.Dial(); err = `%v`; want %v", err, nil)
		return
	}
	defer client.Close()

	var documents []model.Document
	err = client.Call("Server.Search", "Title2", &documents)
	if err != nil {
		t.Errorf("client.Call(); err = `%v`; want %v", err, nil)
	}

	got := documents
	want := []model.Document{
		{ID: 2, URL: "URL 2", Title: "Title2"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("documents = %v; want %v", got, want)
	}
}
