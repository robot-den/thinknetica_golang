package btree

import (
	"pkg/crawler/stubscnr"
	"pkg/index/word"
	"pkg/storage/memory"
	"testing"
)

func TestService_Search(t *testing.T) {
	scanner := stubscnr.New()
	data, _ := scanner.Scan()
	storage, _ := memory.NewStorage()

	ind := word.NewService(storage)
	err := ind.Fill(&data)
	if err != nil {
		t.Errorf("ind.Fill(&data); err = %s; want nil", err)
		return
	}

	eng := NewService(storage)
	found, err := eng.Search("three")
	if err != nil {
		t.Errorf("eng.Search(); err = %s; want nil", err)
		return
	}
	got := len(found)
	want := 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found, err = eng.Search("three.com")
	if err != nil {
		t.Errorf("eng.Search(); err = %s; want nil", err)
		return
	}
	got = len(found)
	want = 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}
