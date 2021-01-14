package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"pkg/auth"
	"pkg/model"
	"strings"
	"testing"
)

var api *API

func TestMain(m *testing.M) {
	auth := auth.New()
	api = New(auth, ":9000")
	api.endpoints()

	os.Exit(m.Run())
}

func TestAPI_Serve_auth(t *testing.T) {
	creds := model.Credentials{
		Login:    "admin",
		Password: "strong",
	}
	encodedCreds, _ := json.Marshal(creds)
	req := httptest.NewRequest(http.MethodPost, AuthPath, bytes.NewReader(encodedCreds))
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}

	got := rec.Body.String()
	// NOTE: the rest part of the token is dynamic
	want := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

	if !strings.Contains(got, want) {
		t.Errorf("strings.Contains(got, want) = %v; want %v", false, true)
	}
}

func TestAPI_Serve_auth_not_found(t *testing.T) {
	creds := model.Credentials{
		Login:    "hacker",
		Password: "password",
	}
	encodedCreds, _ := json.Marshal(creds)
	req := httptest.NewRequest(http.MethodPost, AuthPath, bytes.NewReader(encodedCreds))
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusNotFound)
	}

	got := strings.TrimSpace(rec.Body.String()) // Почему-то http.Error() добавляет в конец строки перенос и пробелы
	want := "пользователь с таким логином или паролем не существует"

	if got != want {
		t.Errorf("rec.Body = %s; want: %s", got, want)
	}
}

func TestAPI_Serve_options(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, RootPath, nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}

	want := "OPTIONS, POST"
	got := rec.Header().Get("Allow")

	if got != want {
		t.Errorf("`Allow` header = %s, want %s", got, want)
	}
}
