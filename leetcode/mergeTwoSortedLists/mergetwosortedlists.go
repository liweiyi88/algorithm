package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	var list, loopList *ListNode

	if list1.Val <= list2.Val {
		list, loopList = list1, list2
	} else {
		list, loopList = list2, list1
	}

	prev := list
	next := list.Next

	for loopList != nil {
		for next != nil && next.Val < loopList.Val {
			prev = next
			next = next.Next
		}

		next2 := loopList.Next
		prev.Next = loopList
		prev = prev.Next
		loopList.Next = next
		loopList = next2
	}

	return list
}
