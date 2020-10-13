package fibonacci

import "testing"

func TestAt(t *testing.T) {
	got := At(10)
	if got != 55 {
		t.Errorf("At(10) = %d; want 55", got)
	}
	got = At(0)
	if got != 0 {
		t.Errorf("At(0) = %d; want 0", got)
	}

	got = At(1)
	if got != 1 {
		t.Errorf("At(1) = %d; want 1", got)
	}
}
