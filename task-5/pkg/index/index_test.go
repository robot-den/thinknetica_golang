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

func TestBTree_Add(t *testing.T) {
	tree := BTree{}

	got := tree.String()
	want := "[]"
	if got != want {
		t.Errorf("tree.String() = %s; want %s", got, want)
	}

	tree.Add(&Record{Id: 5})
	got = tree.String()
	want = "[5]"
	if got != want {
		t.Errorf("tree.String() = %s; want %s", got, want)
	}

	tree.Add(&Record{Id: 3})
	tree.Add(&Record{Id: 8})
	tree.Add(&Record{Id: 4})
	tree.Add(&Record{Id: 1})
	got = tree.String()
	want = "[5 3 1 4 8]"
	if got != want {
		t.Errorf("tree.String() = %s; want %s", got, want)
	}
}

func TestBTree_Search(t *testing.T) {
	tree := BTree{}
	rec := &Record{Id: 4}

	tree.Add(&Record{Id: 5})
	tree.Add(&Record{Id: 3})
	tree.Add(&Record{Id: 8})
	tree.Add(rec)
	tree.Add(&Record{Id: 1})

	got, ok := tree.Search(4)
	if !ok {
		t.Errorf("ok in tree.String() = %v; want true", ok)
	}
	want := rec
	if got != want {
		t.Errorf("tree.String() = %v; want %v", got, want)
	}

	_, ok = tree.Search(7)
	if ok {
		t.Errorf("ok in tree.String() = %v; want false", ok)
	}
}
