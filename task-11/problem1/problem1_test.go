package problem1

import (
	"testing"
)

func TestMaxAge(t *testing.T) {
	c1 := Customer{age: 25}
	c2 := Customer{age: 45}
	e1 := Employee{age: 55}
	e2 := Employee{age: 32}

	want := e1.Age()
	got := MaxAge(c1, c2, e1, e2)
	if got != want {
		t.Errorf("MaxAge() = %v, want %v", got, want)
	}
}
