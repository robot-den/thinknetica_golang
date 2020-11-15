package btree

import (
	"fmt"
	"pkg/crawler/stubscnr"
	"pkg/index/word"
	"pkg/storage/memory"
	"testing"
)

func TestService_Search(t *testing.T) {
	scanner := stubscnr.New()
	data, _ := scanner.Scan()
	storage, _ := memory.NewStorage()
	ind := word.NewService(storage)

	err := ind.Fill(&data)
	if err != nil {
		t.Errorf("ind.Fill(&data); err = %s; want nil", err)
		return
	}

	eng := NewService(storage)
	found, err := eng.Search("three")
	if err != nil {
		t.Errorf("eng.Search(); err = %s; want nil", err)
		return
	}
	got := len(found)
	want := 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found, err = eng.Search("three.com")
	if err != nil {
		t.Errorf("eng.Search(); err = %s; want nil", err)
		return
	}
	got = len(found)
	want = 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}

func BenchmarkService_Search10(b *testing.B) {
	benchmarkService_Search(10, b)
}
func BenchmarkService_Search100(b *testing.B) {
	benchmarkService_Search(100, b)
}
func BenchmarkService_Search1000(b *testing.B) {
	benchmarkService_Search(1000, b)
}
func BenchmarkService_Search10000(b *testing.B) {
	benchmarkService_Search(10000, b)
}
func BenchmarkService_Search50000(b *testing.B) {
	benchmarkService_Search(50000, b)
}

func benchmarkService_Search(n int, b *testing.B) {
	scanner := stubscnr.New()
	data, _ := scanner.ScanN(n)
	storage, _ := memory.NewStorage()
	ind := word.NewService(storage)
	_ = ind.Fill(&data)
	eng := NewService(storage)
	// Try to get something from the middle of generated data
	phrase := fmt.Sprintf("Lorem%d", n/2)
	// Reset timer to remove expensive Fill() setup time from calculations time
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result, _ := eng.Search(phrase)
		_ = result
	}
}
