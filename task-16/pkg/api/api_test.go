package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"pkg/engine"
	"pkg/index/hash"
	"pkg/model"
	"pkg/storage/memory"
	"reflect"
	"testing"
)

var api *Service

func TestMain(m *testing.M) {
	str := memory.NewStorage()
	ind := hash.NewService()
	eng := engine.NewService(ind, str)
	docs := []model.Document{
		{URL: "URL 1", Title: "Title1"},
		{URL: "URL 2", Title: "Title2"},
	}
	eng.BatchCreate(docs)

	api = NewService(eng, ":9000")
	api.endpoints()

	os.Exit(m.Run())
}

func TestService_Serve_documents_index(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, DocumentsPath, nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}

	docs := []model.Document{}
	_ = json.NewDecoder(rec.Body).Decode(&docs)
	got := []int{}
	for _, doc := range docs {
		got = append(got, doc.ID)
	}
	want := []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("reflect.DeepEqual(got, want) = %v; want %v", false, true)
	}
}

func TestService_Serve_create_document(t *testing.T) {
	doc := model.Document{URL: "URL 3", Title: "Title 3"}
	encodedDoc, _ := json.Marshal(doc)

	req := httptest.NewRequest(http.MethodPost, DocumentsPath, bytes.NewReader(encodedDoc))
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusCreated)
	}

	savedDoc := model.Document{}
	_ = json.NewDecoder(rec.Body).Decode(&savedDoc)
	got := savedDoc.ID
	want := 3
	if got != want {
		t.Errorf("savedDoc.ID = %v; want %v", savedDoc.ID, want)
	}
}

func TestService_Serve_get_document(t *testing.T) {
	url := DocumentsPath + "/2"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}

	want := model.Document{ID: 2, URL: "URL 2", Title: "Title2"}
	got := model.Document{}
	_ = json.NewDecoder(rec.Body).Decode(&got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got = %v; want %v", got, want)
	}
}

func TestService_Serve_update_document(t *testing.T) {
	doc := model.Document{ID: 1, URL: "New URL", Title: "New Title"}
	encodedDoc, _ := json.Marshal(doc)

	req := httptest.NewRequest(http.MethodPut, DocumentsPath, bytes.NewReader(encodedDoc))
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}
}

func TestService_Serve_delete_document(t *testing.T) {
	url := DocumentsPath + "/3"
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}
}

func TestService_Serve_search_documents(t *testing.T) {
	url := DocumentsSearchPath + "?token=Title2"

	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}

	foundDocs := []model.Document{}
	_ = json.NewDecoder(rec.Body).Decode(&foundDocs)

	got := []int{}
	for _, doc := range foundDocs {
		got = append(got, doc.ID)
	}
	want := []int{2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("reflect.DeepEqual(got, want) = %v; want %v", false, true)
	}
}

func TestService_Serve_options(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, RootPath, nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}

	want := "OPTIONS, GET, POST, PUT, DELETE"
	got := rec.Header().Get("Allow")

	if got != want {
		t.Errorf("`Allow` header = %s, want %s", got, want)
	}
}
