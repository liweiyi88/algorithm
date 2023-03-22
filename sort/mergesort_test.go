package sort

import "testing"

func TestMergeSort(t *testing.T) {
	example := []string{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	mergesort := MergeSort{}
	mergesort.Sort(example)

	expected := []string{"A", "E", "E", "E", "E", "G", "L", "M", "M", "O", "P", "R", "R", "S", "T", "X"}
	for i, v := range example {
		if v != expected[i] {
			t.Errorf("expected: %v, actual: %v", expected, example)
		}
	}
}
