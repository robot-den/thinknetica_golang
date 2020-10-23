package crawler

import (
	"testing"
)

func TestScanSite(t *testing.T) {
	const url = "https://habr.com"
	const depth = 1
	c := New(url, depth)
	data, err := c.Scan()
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}
