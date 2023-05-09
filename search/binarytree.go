package search

import (
	"fmt"
	"io"
)

type node struct {
	n           int
	key         string
	value       any
	left, right *node
}

func newNode(key string, value any, n int) *node {
	return &node{
		n:     n,
		key:   key,
		value: value,
	}
}

type BinaryTree struct {
	root *node
}

func delete(x *node, key string) *node {
	if x == nil {
		return nil
	}

	if key < x.key {
		x.left = delete(x.left, key)
	} else if key > x.key {
		x.right = delete(x.right, key)
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

func deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}

	x.left = deleteMin(x.left)
	x.n = size(x.left) + size(x.right) + 1
	return x
}

func get(x *node, key string) any {
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

func min(x *node) *node {
	if x.left == nil {
		return x
	}

	return min(x.left)
}

func put(x *node, key string, value any) *node {
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

func size(x *node) int {
	if x == nil {
		return 0
	}

	return x.n
}

func (b *BinaryTree) Get(key string) any {
	return get(b.root, key)
}

func (b *BinaryTree) Put(key string, value any) {
	b.root = put(b.root, key, value)
}

func (b *BinaryTree) Size() int {
	return size(b.root)
}

// minimum key
func (b *BinaryTree) Min() string {
	return min(b.root).key
}
func (b *BinaryTree) DeleteMin() {
	b.root = deleteMin(b.root)
}

func (b *BinaryTree) Delete(key string) {
	b.root = delete(b.root, key)
}

func print(x *node, w io.Writer) {
	if x.left != nil {
		print(x.left, w)
	}

	fmt.Fprint(w, x.key)

	if x.right != nil {
		print(x.right, w)
	}
}

func (b *BinaryTree) Print(w io.Writer) {
	print(b.root, w)
}
