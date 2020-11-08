package btree

import (
	"pkg/crawler/stub"
	"pkg/index/word"
	"pkg/storage/memory"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := stub.New()
	data, _ := scanner.Scan()
	storage, err := memory.NewStorage()
	if err != nil {
		t.Errorf("memory.NewStorage(); err = %s; want nil", err)
		return
	}
	ind := word.NewService(storage)
	err = ind.Fill(&data)
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
