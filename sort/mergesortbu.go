package sort

import (
	"golang.org/x/exp/constraints"
)

// Merge sort bottom up approach
type MergeSortBU[K constraints.Ordered] struct {
}

func (m *MergeSortBU[K]) Sort(a []K) {
	aux := make([]K, len(a))
	m.sort(a, aux, 0, len(a)-1)
}

func (m *MergeSortBU[K]) sort(a, aux []K, low, high int) {
	if high <= low {
		return
	}

	for i := 1; i < len(a); i *= 2 {
		for j := 0; j < len(a)-i; j += i + i {
			m.merge(a, aux, j, i+j-1, min(j+i+i-1, len(a)-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m *MergeSortBU[K]) merge(a, aux []K, low, mid, high int) {
	i, j := low, mid+1

	for k := low; k <= high; k++ {
		aux[k] = a[k]
	}

	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > high {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
