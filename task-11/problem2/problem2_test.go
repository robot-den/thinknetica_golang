package problem2

import (
	"testing"
)

func TestMaxAge(t *testing.T) {
	c1 := Customer{Age: 25}
	c2 := Customer{Age: 45}
	e1 := Employee{Age: 55}
	e2 := Employee{Age: 32}

	var want interface{} = e1
	got, _ := MaxAge(c1, c2, e1, e2)
	if got != want {
		t.Errorf("MaxAge() = %v, want %v", got, want)
	}
}
