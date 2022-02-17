package easy

// ListNode is a Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	output := &ListNode{}
	// have to copy the output pointer to build it without incrementing it
	nextNode := output
	nextNodeExists := true
	for nextNodeExists {
		nextNodeExists = false // false until proven true

		sumOfNodes := l1.Val + l2.Val + carry
		carry = 0
		if sumOfNodes > 9 {
			carry = (sumOfNodes - (sumOfNodes % 10)) / 10
		}
		nextNode.Next = &ListNode{ Val:  sumOfNodes % 10, Next: nil}
		nextNode = nextNode.Next

		if l1.Next != nil || l2.Next != nil || carry != 0 {
            // ^ proof that sequence needs to continue
			nextNodeExists = true
			if l1.Next == nil {
				l1.Next = &ListNode{Val: 0, Next: nil}
			}
			if l2.Next == nil {
				l2.Next = &ListNode{Val: 0, Next: nil}
			}
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return output.Next
}
