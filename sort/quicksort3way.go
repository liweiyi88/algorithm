package sort

import "golang.org/x/exp/constraints"

type Quick3Way[T constraints.Ordered] struct{}

func (qs *Quick3Way[T]) Sort(a []T) {
	qs.sort(a, 0, len(a)-1)
}

func (qs *Quick3Way[T]) sort(a []T, low, high int) {
	if high <= low {
		return
	}

	lt, i, gt := low, low+1, high
	v := a[low]

	for i <= gt {
		if a[i] < v {
			a[i], a[lt] = a[lt], a[i]
			i++
			lt++
		} else if a[i] > v {
			a[i], a[gt] = a[gt], a[i]
			gt--
		} else {
			i++
		}
	}

	qs.sort(a, low, lt-1)
	qs.sort(a, gt+1, high)
}
