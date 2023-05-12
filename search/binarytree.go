package search

import (
	"fmt"
	"io"

	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	n           int
	key         T
	value       any
	left, right *node[T]
}

func newNode[T constraints.Ordered](key T, value any, n int) *node[T] {
	return &node[T]{
		n:     n,
		key:   key,
		value: value,
	}
}

type BinaryTree[T constraints.Ordered] struct {
	root *node[T]
}

func delete[T constraints.Ordered](x *node[T], key T) *node[T] {
	if x == nil {
		return nil
	}

	if key < x.key {
		x.left = delete[T](x.left, key)
	} else if key > x.key {
		x.right = delete[T](x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}

		if x.left == nil {
			return x.right
		}

		t := x
		x = min(t.right)
		x.right = deleteMin(t.right)
		x.left = t.left
	}

	x.n = size(x.left) + size(x.right) + 1

	return x
}

func deleteMin[T constraints.Ordered](x *node[T]) *node[T] {
	if x.left == nil {
		return x.right
	}

	x.left = deleteMin(x.left)
	x.n = size(x.left) + size(x.right) + 1
	return x
}

func get[T constraints.Ordered](x *node[T], key T) any {
	if x == nil {
		return nil
	}

	if key < x.key {
		return get(x.left, key)
	}

	if key > x.key {
		return get(x.right, key)
	}

	return x.value
}

func keys[T constraints.Ordered](x *node[T], queue *[]T, lowKey, highKey T) {
	if x == nil {
		return
	}

	if lowKey < x.key {
		keys(x.left, queue, lowKey, highKey)
	}

	if lowKey <= x.key && highKey >= x.key {
		*queue = append(*queue, x.key)
	}

	if highKey > x.key {
		keys(x.right, queue, lowKey, highKey)
	}
}

func min[T constraints.Ordered](x *node[T]) *node[T] {
	if x.left == nil {
		return x
	}

	return min(x.left)
}

func print[T constraints.Ordered](x *node[T], w io.Writer) {
	if x.left != nil {
		print(x.left, w)
	}

	fmt.Fprint(w, x.key)

	if x.right != nil {
		print(x.right, w)
	}
}

func put[T constraints.Ordered](x *node[T], key T, value any) *node[T] {
	if x == nil {
		return newNode(key, value, 1)
	}

	if key < x.key {
		x.left = put(x.left, key, value)
	} else if key > x.key {
		x.right = put(x.right, key, value)
	} else {
		x.value = value
	}

	x.n = size(x.left) + size(x.right) + 1

	return x
}

func rank[T constraints.Ordered](key T, x *node[T]) int {
	if x == nil {
		return 0
	}

	if key == x.key {
		return size(x.left)
	}

	if key < x.key {
		return rank(key, x.left)
	}

	return 1 + size(x.left) + rank(key, x.right)
}

func size[T constraints.Ordered](x *node[T]) int {
	if x == nil {
		return 0
	}

	return x.n
}

func (b *BinaryTree[T]) Get(key T) any {
	return get(b.root, key)
}

func (b *BinaryTree[T]) Put(key T, value any) {
	b.root = put(b.root, key, value)
}

func (b *BinaryTree[T]) Size() int {
	return size(b.root)
}

// minimum key
func (b *BinaryTree[T]) Min() T {
	return min(b.root).key
}
func (b *BinaryTree[T]) DeleteMin() {
	b.root = deleteMin(b.root)
}

func (b *BinaryTree[T]) Delete(key T) {
	b.root = delete(b.root, key)
}

func (b *BinaryTree[T]) Keys(lowKey, highKey T) []T {
	queue := make([]T, 0, b.Size())
	keys(b.root, &queue, lowKey, highKey)

	return queue
}

func (b *BinaryTree[T]) Print(w io.Writer) {
	print(b.root, w)
}

func (b *BinaryTree[T]) Rank(key T) int {
	return rank(key, b.root)
}
