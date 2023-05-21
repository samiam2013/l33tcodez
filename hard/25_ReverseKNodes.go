package hard

func reverseKGroup(head *ListNode, k int) *ListNode {
	// if there are less than k nodes, return the head
	numNodes := 0
	for node := head; node != nil; node = node.Next {
		numNodes++
	}
	if numNodes < k {
		return head
	}

	// reverse the first k nodes
	var prev, next *ListNode
	node := head
	for i := 0; i < k; i++ {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}

	// recursively reverse the rest of the list
	head.Next = reverseKGroup(next, k)
	return prev

}
