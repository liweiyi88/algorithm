package trappingrainwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	assert.Equal(t, 0, trap([]int{0, 1, 1, 1, 1, 1, 2, 1}))
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))
	assert.Equal(t, 1, trap([]int{4, 2, 3}))
	assert.Equal(t, 1, trap([]int{5, 4, 1, 2}))
	assert.Equal(t, 0, trap([]int{4, 4, 4, 7, 1, 0}))
	assert.Equal(t, 7, trap([]int{0, 7, 1, 4, 6}))
	assert.Equal(t, 1, trap([]int{4, 9, 4, 5, 3, 2}))
	assert.Equal(t, 0, trap([]int{0}))
	assert.Equal(t, 0, trap([]int{0, 2, 0}))
	assert.Equal(t, 26, trap([]int{0, 1, 2, 0, 3, 0, 1, 2, 0, 0, 4, 2, 1, 2, 5, 0, 1, 2, 0, 2}))
	assert.Equal(t, 12, trap([]int{2, 8, 5, 5, 6, 1, 7, 4, 5}))
}
