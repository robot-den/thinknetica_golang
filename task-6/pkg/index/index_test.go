package index

import (
	"pkg/crawler"
	"pkg/storage/memory"
	"testing"
)

func TestFill(t *testing.T) {
	scanner := crawler.NewStub()
	data, _ := scanner.Scan()
	storage, err := memory.NewStorage()
	if err != nil {
		t.Errorf("memory.NewStorage(); err = %s; want nil", err)
		return
	}
	ind := New(storage)
	err = ind.Fill(&data)
	if err != nil {
		t.Errorf("ind.Fill(); err = %s; want nil", err)
		return
	}

	records, index, err := storage.Read()
	if err != nil {
		t.Errorf("storage.Read(); err = %s; want nil", err)
		return
	}
	got := len(records)
	want := 3
	if got != want {
		t.Errorf("len(records) = %d; want %d", got, want)
	}
	{
		_, got := index["three.com"]
		want := true
		if got != want {
			t.Errorf("index[\"three\"]; got ok = %t; want %t", got, want)
		}
	}
}
