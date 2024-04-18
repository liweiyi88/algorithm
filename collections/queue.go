package collections

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var v T
		return v, false
	}

	item := (*q)[0]
	*q = (*q)[1:len(*q)]

	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}
