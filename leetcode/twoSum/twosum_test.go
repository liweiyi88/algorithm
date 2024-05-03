package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{0, 2}, twoSum([]int{3, 2, 3}, 6))
	assert.Equal(t, []int{0, 3}, twoSum([]int{0, 4, 3, 0}, 0))
}
