package webapp

import (
	"net/http"
	"net/http/httptest"
	"pkg/index/hash"
	"pkg/storage/memory"
	"strings"
	"testing"
)

func TestWebApp_Run(t *testing.T) {
	str := memory.NewStorage()
	ind := hash.NewService()
	wa := New(ind, str, ":9000")
	router := wa.endpoints()

	tests := []struct {
		name  string
		route string
		title string
	}{
		{name: "Root", route: "/", title: "Links"},
		{name: "Index", route: "/index", title: "Index"},
		{name: "Docs", route: "/docs", title: "Docs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.route, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
			}
			if !strings.Contains(rec.Body.String(), tt.title) {
				t.Errorf("strings.Contains(rec.Body.String(), \"%s\") = %v; want %v", tt.title, false, true)
			}
		})
	}
}
