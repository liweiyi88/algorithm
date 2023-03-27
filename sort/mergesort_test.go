package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	example := []string{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	mergesort := MergeSort[string]{}
	mergesort.Sort(example)

	expected := []string{"A", "E", "E", "E", "E", "G", "L", "M", "M", "O", "P", "R", "R", "S", "T", "X"}
	for i, v := range example {
		if v != expected[i] {
			t.Errorf("expected: %v, actual: %v", expected, example)
		}
	}
}

func BenchmarkMergeSortLen10(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		sample := rand.Perm(10)
		mergesort := MergeSort[int]{}
		mergesort.Sort(sample)
	}
}

func BenchmarkMergeSortLenMillion(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		sample := rand.Perm(1000000)
		mergesort := MergeSort[int]{}
		mergesort.Sort(sample)
	}
}
