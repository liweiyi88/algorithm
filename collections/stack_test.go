package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := Stack[int]{}

	stack.Push(2)
	stack.Push(3)

	item, _ := stack.Pop()

	assert.Equal(t, 3, item)
	assert.Equal(t, 1, stack.Size())

	stack.Pop()

	assert.True(t, stack.IsEmpty())

	_, ok := stack.Pop()

	assert.False(t, ok)
}
