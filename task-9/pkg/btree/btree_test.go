package btree

import (
	"pkg/model"
	"testing"
)

func TestBTree_Add(t *testing.T) {
	tree := BTree{}

	tree.Add(&model.Document{ID: 5})
	tree.Add(&model.Document{ID: 3})
	tree.Add(&model.Document{ID: 8})
	tree.Add(&model.Document{ID: 4})
	tree.Add(&model.Document{ID: 1})
	got := tree.String()
	want := "[5 3 1 4 8]"
	if got != want {
		t.Errorf("tree.String() = %s; want %s", got, want)
	}
}

func TestBTree_Search(t *testing.T) {
	tree := BTree{}
	doc := &model.Document{ID: 4}

	tree.Add(&model.Document{ID: 5})
	tree.Add(&model.Document{ID: 3})
	tree.Add(&model.Document{ID: 8})
	tree.Add(doc)
	tree.Add(&model.Document{ID: 1})

	got, ok := tree.Search(4)
	if !ok {
		t.Errorf("ok in tree.String() = %v; want true", ok)
	}
	want := doc
	if got != want {
		t.Errorf("tree.String() = %v; want %v", got, want)
	}

	_, ok = tree.Search(7)
	if ok {
		t.Errorf("ok in tree.String() = %v; want false", ok)
	}
}
