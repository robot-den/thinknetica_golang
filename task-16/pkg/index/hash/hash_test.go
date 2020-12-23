package hash

import (
	"pkg/model"
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {
	index := NewService()
	docs := []model.Document{
		{ID: 1, URL: "http://a.com/", Title: "AAA"},
		{ID: 2, URL: "http://b.com/", Title: "BBB"},
		{ID: 3, URL: "http://c.com/", Title: "CCC"},
	}
	index.Update(docs)

	got := index.Search("BBB")
	want := []int{2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("index.Search() = %v; want %v", got, want)
	}
}
