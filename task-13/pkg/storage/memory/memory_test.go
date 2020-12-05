package memory

import (
	"pkg/model"
	"reflect"
	"testing"
)

func TestStorage_Write(t *testing.T) {
	storage := NewStorage()
	rawDocs := []model.Document{
		{ID: 1, URL: "http://a.com/", Title: "AAA"},
		{ID: 2, URL: "http://b.com/", Title: "BBB"},
		{ID: 3, URL: "http://c.com/", Title: "CCC"},
	}
	docs := storage.Write(rawDocs)

	var ids []int
	for _, doc := range docs {
		ids = append(ids, doc.ID)
	}

	got := ids
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("docs IDs after save = %v; want %v", got, want)
	}
}

func TestStorage_Read(t *testing.T) {
	storage := NewStorage()
	rawDocs := []model.Document{
		{ID: 1, URL: "http://a.com/", Title: "AAA"},
		{ID: 2, URL: "http://b.com/", Title: "BBB"},
		{ID: 3, URL: "http://c.com/", Title: "CCC"},
	}
	docs := storage.Write(rawDocs)

	want := "BBB"
	var id int
	for _, doc := range docs {
		if doc.Title == want {
			id = doc.ID
			break
		}
	}

	var got string
	readDocs := storage.Read([]int{id})
	for _, doc := range readDocs {
		got = doc.Title
	}

	if got != want {
		t.Errorf("required doc title = %s; want %s", got, want)
	}

	{
		want := 1
		got := len(readDocs)
		if got != want {
			t.Errorf("len(readDocs) = %d; want %d", got, want)
		}
	}
}
