package engine

import (
	"pkg/stub"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := stub.NewScanner()

	found, err := Search(scanner, "little")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got := len(found)
	want := 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found, err = Search(scanner, "the")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got = len(found)
	want = 2
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found, err = Search(scanner, "doesn't exist")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got = len(found)
	want = 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}
