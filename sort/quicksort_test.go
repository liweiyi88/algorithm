package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	example := []string{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	quickSort := QuickSort[string]{}
	quickSort.Sort(example)

	expected := []string{"A", "E", "E", "E", "E", "G", "L", "M", "M", "O", "P", "R", "R", "S", "T", "X"}
	for i, v := range example {
		if v != expected[i] {
			t.Errorf("expected: %v, actual: %v", expected, example)
		}
	}
}

func BenchmarkQuickSortLen10(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		sample := rand.Perm(10)
		qs := QuickSort[int]{}
		qs.Sort(sample)
	}
}

func BenchmarkQuickSortLenMillion(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		sample := rand.Perm(1000000)
		qs := QuickSort[int]{}
		qs.Sort(sample)
	}
}
