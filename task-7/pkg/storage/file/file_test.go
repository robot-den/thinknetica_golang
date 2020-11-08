package file

import (
	"fmt"
	"os"
	"pkg/model"
	"testing"
)

func TestWrite_Read(t *testing.T) {
	// Clean the state. We ignore possible errors because they all are *PathError type
	_ = os.Remove(DocumentsFile)
	_ = os.Remove(IndexFile)

	store, err := NewStorage()
	if err != nil {
		t.Errorf("NewStorage(); err = %s; want nil", err)
		return
	}

	documents := []model.Document{
		model.Document{
			ID:    1,
			URL:   "http://one.com",
			Title: "First title",
		},
		model.Document{
			ID:    2,
			URL:   "http://two.com",
			Title: "Second title",
		},
	}
	index := model.InvertedIndex{}
	index["one"] = []int{1, 3}
	index["two"] = []int{1, 2}

	err = store.Write(documents, index)
	if err != nil {
		t.Errorf("store.Write(); err = %s; want nil", err)
		return
	}

	readDocuments, readIndex, err := store.Read()
	if err != nil {
		t.Errorf("store.Read(); err = %s; want nil", err)
		return
	}
	want := documents[0]
	got := readDocuments[0]
	if got != want {
		t.Errorf("readDocuments[0] = %v; want %v", got, want)
	}
	want = documents[1]
	got = readDocuments[1]
	if got != want {
		t.Errorf("readDocuments[1] = %v; want %v", got, want)
	}
	{
		want := fmt.Sprint(index["one"])
		got := fmt.Sprint(readIndex["one"])
		if got != want {
			t.Errorf("readIndex[\"one\"] = %s; want %s", got, want)
		}
		want = fmt.Sprint(index["two"])
		got = fmt.Sprint(readIndex["two"])
		if got != want {
			t.Errorf("readIndex[\"two\"] = %s; want %s", got, want)
		}
	}

	documents = []model.Document{
		model.Document{
			ID:    3,
			URL:   "http://three.com",
			Title: "Third title",
		},
	}
	index = model.InvertedIndex{}
	index["three"] = []int{4, 5}

	err = store.Write(documents, index)
	if err != nil {
		t.Errorf("store.Write(); err = %s; want nil", err)
		return
	}
	readDocuments, readIndex, err = store.Read()
	if err != nil {
		t.Errorf("store.Read(); err = %s; want nil", err)
		return
	}
	want = documents[0]
	got = readDocuments[0]
	if got != want {
		t.Errorf("readDocuments[0] = %v; want %v", got, want)
	}
	{
		want := fmt.Sprint(index["three"])
		got := fmt.Sprint(readIndex["three"])
		if got != want {
			t.Errorf("readIndex[\"three\"] = %s; want %s", got, want)
		}
	}

	_ = os.Remove(DocumentsFile)
	_ = os.Remove(IndexFile)
}
