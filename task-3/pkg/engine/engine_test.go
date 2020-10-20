package engine

import (
	"pkg/crawler_stub"
	"testing"
)

func TestSearch(t *testing.T) {
	stub := crawlerStub.New()

	found, err := Search(stub, "little")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got := len(found)
	want := 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found, err = Search(stub, "the")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got = len(found)
	want = 2
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found, err = Search(stub, "doesn't exist")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got = len(found)
	want = 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}
