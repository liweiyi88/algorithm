package sort

import (
	"golang.org/x/exp/constraints"
)

type MergeSort[K constraints.Ordered] struct {
	aux []K
}

func (m *MergeSort[K]) Sort(a []K) {
	m.aux = make([]K, len(a))
	m.sort(a, 0, len(a)-1)
}

func (m *MergeSort[K]) sort(a []K, low, high int) {
	if high <= low {
		return
	}

	mid := low + (high-low)/2
	m.sort(a, low, mid)
	m.sort(a, mid+1, high)
	m.merge(a, low, mid, high)
}

func (m *MergeSort[K]) merge(a []K, low, mid, high int) {
	i, j := low, mid+1

	for k := low; k <= high; k++ {
		m.aux[k] = a[k]
	}

	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = m.aux[j]
			j++
		} else if j > high {
			a[k] = m.aux[i]
			i++
		} else if m.aux[j] < m.aux[i] {
			a[k] = m.aux[j]
			j++
		} else {
			a[k] = m.aux[i]
			i++
		}
	}
}
