package problem3

import (
	"bytes"
	"testing"
)

func Test_write(t *testing.T) {
	v1 := "String"
	v2 := true
	v3 := 10.5
	v4 := struct{ age int }{age: 32}
	v5 := "String2"

	want := "StringString2"

	var buf bytes.Buffer
	write(&buf, v1, v2, v3, v4, v5)
	got := buf.String()
	if got != want {
		t.Errorf("buf.String() = %v, want %v", got, want)
	}
}
