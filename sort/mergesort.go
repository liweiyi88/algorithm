package sort

import (
	"golang.org/x/exp/constraints"
)

type MergeSort[K constraints.Ordered] struct {
}

func (m *MergeSort[K]) Sort(a []K) {
	aux := make([]K, len(a))
	m.sort(a, aux, 0, len(a)-1)
}

func (m *MergeSort[K]) sort(a, aux []K, low, high int) {
	if high <= low {
		return
	}

	mid := low + (high-low)/2
	m.sort(a, aux, low, mid)
	m.sort(a, aux, mid+1, high)
	m.merge(a, aux, low, mid, high)
}

func (m *MergeSort[K]) merge(a, aux []K, low, mid, high int) {
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
