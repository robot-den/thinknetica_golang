package memory

import (
	"fmt"
	"pkg/model"
	"testing"
)

var documentsSet1 = []model.Document{
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

var index1 = model.InvertedIndex{
	"one": []int{1, 3},
	"two": []int{1, 2},
}

func TestStorage_Write(t *testing.T) {
	store, _ := NewStorage()

	err := store.Write(documentsSet1, index1)
	if err != nil {
		t.Errorf("store.Write(); err = %s; want nil", err)
		return
	}
	readDocuments, readIndex, err := store.Read()
	if err != nil {
		t.Errorf("store.Read(); err = %s; want nil", err)
		return
	}
	want := documentsSet1[0]
	got := readDocuments[0]
	if got != want {
		t.Errorf("readDocuments[0] = %v; want %v", got, want)
	}
	want = documentsSet1[1]
	got = readDocuments[1]
	if got != want {
		t.Errorf("readDocuments[1] = %v; want %v", got, want)
	}
	{
		want := fmt.Sprint(index1["one"])
		got := fmt.Sprint(readIndex["one"])
		if got != want {
			t.Errorf("readIndex[\"one\"] = %s; want %s", got, want)
		}
		want = fmt.Sprint(index1["two"])
		got = fmt.Sprint(readIndex["two"])
		if got != want {
			t.Errorf("readIndex[\"two\"] = %s; want %s", got, want)
		}
	}
}

func TestStorage_Write_rewrite(t *testing.T) {
	store, _ := NewStorage()

	_ = store.Write(documentsSet1, index1)

	documentsSet2 := []model.Document{
		model.Document{
			ID:    3,
			URL:   "http://three.com",
			Title: "Third title",
		},
	}
	index2 := model.InvertedIndex{
		"three": []int{4, 5},
	}

	err := store.Write(documentsSet2, index2)
	if err != nil {
		t.Errorf("store.Write(); err = %s; want nil", err)
		return
	}

	readDocuments, readIndex, err := store.Read()
	if err != nil {
		t.Errorf("store.Read(); err = %s; want nil", err)
		return
	}
	want := documentsSet2[0]
	got := readDocuments[0]
	if got != want {
		t.Errorf("readDocuments[0] = %v; want %v", got, want)
	}
	{
		want := fmt.Sprint(index2["three"])
		got := fmt.Sprint(readIndex["three"])
		if got != want {
			t.Errorf("readIndex[\"three\"] = %s; want %s", got, want)
		}
	}
}
