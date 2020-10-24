package index

import (
	"pkg/stub"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := stub.NewScanner()
	data, _ := scanner.Scan()
	ind := New()
	ind.Fill(&data)

	found := ind.Search("http")
	got := len(found)
	want := 3
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found = ind.Search("little")
	got = len(found)
	want = 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found = ind.Search("golang")
	got = len(found)
	want = 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}
