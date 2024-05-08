package mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert := assert.New(t)

	list1 := &ListNode{
		Val: 2,
	}

	list2 := &ListNode{
		Val: 1,
	}

	list := mergeTwoLists(list1, list2)

	var result []int

	for list != nil {
		result = append(result, list.Val)
		list = list.Next
	}

	assert.Equal([]int{1, 2}, result)
}
