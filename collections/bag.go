package collections

type Bag[T comparable] map[T]int

func NewBag[T comparable]() Bag[T] {
	return make(map[T]int)
}

func (bag Bag[T]) Add(item T) {
	bag[item]++
}

func (bag Bag[T]) Count(item T) int {
	v, ok := bag[item]

	if !ok {
		return 0
	}

	return v
}

func (bag Bag[T]) Delete(item T) {
	delete(bag, item)
}

func (bag Bag[T]) IsEmpty() bool {
	return len(bag) == 0
}

func (bag Bag[T]) Size() int {
	return len(bag)
}
