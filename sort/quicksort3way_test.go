package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuick3Way(t *testing.T) {
	example := []string{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	quickSort := Quick3Way[string]{}
	quickSort.Sort(example)

	expected := []string{"A", "E", "E", "E", "E", "G", "L", "M", "M", "O", "P", "R", "R", "S", "T", "X"}
	for i, v := range example {
		if v != expected[i] {
			t.Errorf("expected: %v, actual: %v", expected, example)
		}
	}
}

func BenchmarkQuick3WayLen10(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		sample := rand.Perm(10)
		qs := Quick3Way[int]{}
		qs.Sort(sample)
	}
}

func BenchmarkQuick3WayLenMillion(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		sample := rand.Perm(1000000)
		qs := Quick3Way[int]{}
		qs.Sort(sample)
	}
}
