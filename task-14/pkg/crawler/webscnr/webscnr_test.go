package webscnr

import (
	"testing"
)

func TestWebScnr_Scan(t *testing.T) {
	if !testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	const url = "https://habr.com"
	const depth = 1
	c := WebScnr{}
	data, err := c.Scan(url, depth)
	if err != nil {
		t.Errorf("c.Scan(); err = %s; want nil", err)
		return
	}

	got := len(data)
	want := 1
	if got != want {
		t.Errorf("len(data) = %d; want %d", got, want)
	}
}
