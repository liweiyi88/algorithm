package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := Queue[int]{}

	queue.Enqueue(2)
	queue.Enqueue(3)

	item, _ := queue.Dequeue()

	assert.Equal(t, 2, item)
	assert.Equal(t, 1, queue.Size())

	queue.Dequeue()

	assert.True(t, queue.IsEmpty())

	_, ok := queue.Dequeue()

	assert.False(t, ok)
}
