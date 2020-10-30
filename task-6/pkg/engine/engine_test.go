package engine

import (
	"pkg/crawler"
	"pkg/index"
	"pkg/storage/memory"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := crawler.NewStub()
	data, _ := scanner.Scan()
	storage, err := memory.NewStorage()
	if err != nil {
		t.Errorf("memory.NewStorage(); err = %s; want nil", err)
		return
	}
	ind := index.New(storage)
	err = ind.Fill(&data)
	if err != nil {
		t.Errorf("ind.Fill(&data); err = %s; want nil", err)
		return
	}

	eng := New(storage)
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
