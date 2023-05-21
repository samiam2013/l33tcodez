package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return result
}

func mergeTwoLists(list, list2 *ListNode) *ListNode {
	if list == nil {
		return list2
	} else if list2 == nil {
		return list
	}

	var head *ListNode
	if list.Val < list2.Val {
		head = list
		list = list.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	curNode := head
	for list != nil && list2 != nil {
		if list.Val <= list2.Val {
			curNode.Next = list
			list = list.Next
		} else {
			curNode.Next = list2
			list2 = list2.Next
		}
		curNode = curNode.Next
	}

	if list != nil {
		curNode.Next = list
	} else if list2 != nil {
		curNode.Next = list2
	}

	return head
}
