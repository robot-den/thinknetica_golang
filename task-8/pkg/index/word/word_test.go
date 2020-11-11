package word

import (
	"pkg/crawler/stubscnr"
	"pkg/storage/memory"
	"testing"
)

func TestService_Fill(t *testing.T) {
	scanner := stubscnr.New()
	data, _ := scanner.Scan()
	storage, _ := memory.NewStorage()
	ind := NewService(storage)
	err := ind.Fill(&data)
	if err != nil {
		t.Errorf("ind.Fill(); err = %s; want nil", err)
	}
}

func TestService_storage_integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	scanner := stubscnr.New()
	data, _ := scanner.Scan()
	storage, _ := memory.NewStorage()
	ind := NewService(storage)
	_ = ind.Fill(&data)

	documents, index, err := storage.Read()
	if err != nil {
		t.Errorf("storage.Read(); err = %s; want nil", err)
		return
	}
	got := len(documents)
	want := 3
	if got != want {
		t.Errorf("len(documents) = %d; want %d", got, want)
	}
	{
		_, got := index["three.com"]
		want := true
		if got != want {
			t.Errorf("index[\"three\"]; got ok = %t; want %t", got, want)
		}
	}
}
