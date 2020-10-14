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

// Is this better?
// func TestAt(t *testing.T) {
// 	testSuccess(0, 0, t)
// 	testSuccess(1, 1, t)
// 	testSuccess(10, 55, t)
// 	testSuccess(20, 6765, t)
//
// 	testFail(-1, t)
// 	testFail(21, t)
// }
//
// func testSuccess(n, want int, t *testing.T) {
// 	got, err := At(n)
// 	if err != nil {
// 		t.Errorf("At(%d) failed with error: %s", n, err)
// 	}
// 	if got != want {
// 		t.Errorf("At(%d) = %d; want %d", n, got, want)
// 	}
// }
//
// func testFail(n int, t *testing.T) {
// 	got, err := At(n)
// 	want := 0
// 	if err == nil {
// 		t.Errorf("At(%d) finished successfully but must fail", n)
// 	}
// 	if got != want {
// 		t.Errorf("At(%d) = %d; want %d", n, got, want)
// 	}
// }
