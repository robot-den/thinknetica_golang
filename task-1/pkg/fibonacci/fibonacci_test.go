package fibonacci

import "testing"

func TestAt(t *testing.T) {
	got, _ := At(10)
	want := 55
	if got != want {
		t.Errorf("At(10) = %d; want %d", got, want)
	}
	got, _ = At(0)
	want = 0
	if got != want {
		t.Errorf("At(0) = %d; want %d", got, want)
	}

	got, _ = At(1)
	want = 1
	if got != want {
		t.Errorf("At(1) = %d; want %d", got, want)
	}
}
