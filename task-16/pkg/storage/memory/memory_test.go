package memory

import (
	"pkg/model"
	"reflect"
	"testing"
)

func TestStorage_Write(t *testing.T) {
	storage := NewStorage()
	rawDocs := []model.Document{
		{ID: 0, URL: "http://a.com/", Title: "AAA"},
		{ID: 0, URL: "http://b.com/", Title: "BBB"},
		{ID: 0, URL: "http://c.com/", Title: "CCC"},
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

func TestStorage_Update(t *testing.T) {
	storage := NewStorage()
	doc := model.Document{ID: 0, URL: "http://a.com/", Title: "AAA"}
	savedDoc := storage.Write([]model.Document{doc})[0]

	savedDoc.Title = "BBB"
	err := storage.Update(savedDoc)
	if err != nil {
		t.Errorf("storage.Update(savedDoc); err = %v; want %v", err, nil)
	}

	got := storage.Read([]int{savedDoc.ID})[0]
	want := model.Document{ID: 1, URL: "http://a.com/", Title: "BBB"}
	if got != want {
		t.Errorf("updatedDoc = %v; want %v", got, want)
	}
}

func TestStorage_Update_not_found(t *testing.T) {
	storage := NewStorage()
	notExistedDoc := model.Document{ID: 1, URL: "http://a.com/", Title: "BBB"}

	err := storage.Update(notExistedDoc)
	if err == nil {
		t.Errorf("storage.Update(notExistedDoc); err = %v; want %v", nil, err)
	}
}

func TestStorage_Delete(t *testing.T) {
	storage := NewStorage()
	rawDocs := []model.Document{
		{ID: 0, URL: "http://a.com/", Title: "AAA"},
		{ID: 0, URL: "http://b.com/", Title: "BBB"},
		{ID: 0, URL: "http://c.com/", Title: "CCC"},
	}
	secondDoc := storage.Write(rawDocs)[1]

	err := storage.Delete(secondDoc.ID)
	if err != nil {
		t.Errorf("storage.Delete(secondDoc.ID); err = %v; want %v", err, nil)
	}

	docs := storage.ReadAll()
	var ids []int
	for _, doc := range docs {
		ids = append(ids, doc.ID)
	}
	got := ids
	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("documents ids after delete = %v; want %v", got, want)
	}
}

func TestStorage_Delete_not_found(t *testing.T) {
	storage := NewStorage()
	err := storage.Delete(0)
	if err == nil {
		t.Errorf("storage.Delete(0); err = %v; want %v", nil, err)
	}
}
