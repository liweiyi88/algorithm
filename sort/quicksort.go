package sort

import (
	"golang.org/x/exp/constraints"
)

type QuickSort[T constraints.Ordered] struct {
}

func (qs *QuickSort[T]) Sort(a []T) {
	// Eliminate dependence on input, but it downgrades BenchMark test as we feed random number sample data already.
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	qs.sort(a, 0, len(a)-1)
}

func (qs *QuickSort[T]) partition(a []T, low, high int) int {
	v := a[low]
	i, j := low+1, high

	for {
		for ; i <= high && a[i] <= v; i++ {
		}

		for ; j > low && a[j] >= v; j-- {
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
	}

	a[low], a[j] = a[j], a[low]
	return j
}

func (qs *QuickSort[T]) sort(a []T, low, high int) {
	if high <= low {
		return
	}

	j := qs.partition(a, low, high)
	qs.sort(a, low, j-1)
	qs.sort(a, j+1, high)
}
