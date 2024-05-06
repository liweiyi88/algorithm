package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	input1 := []string{"flower", "flow", "flight"}
	input2 := []string{"dog", "racecar", "car"}
	input3 := []string{"aa", "aa"}
	input4 := []string{"aa", "ab"}

	assert := assert.New(t)
	assert.Equal("fl", longestCommonPrefix(input1))
	assert.Equal("", longestCommonPrefix(input2))
	assert.Equal("aa", longestCommonPrefix(input3))
	assert.Equal("a", longestCommonPrefix(input4))
}
