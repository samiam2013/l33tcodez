package easy

// ListNode is a Definition for singly-linked list. */
type ListNode struct {
	Val int
	Next *ListNode
}

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    nextNodeExists := true
    carry := 0
    output := &ListNode{}
    nextNode := output
    for nextNodeExists {
        sumOfNodes := l1.Val + l2.Val + carry
        if sumOfNodes > 9 {
            sumLeftover := sumOfNodes % 10
            carry = (sumOfNodes - sumLeftover)/10
            sumOfNodes = sumLeftover
        } else {
            carry = 0
        }
        newNextNode := &ListNode{
            Val: sumOfNodes,
            Next: nil,
        }
        nextNode.Next = newNextNode
        nextNode = newNextNode
        
        nextNodeExists = false
        if l1.Next != nil || l2.Next != nil || carry != 0 {
            nextNodeExists = true
            if l1.Next != nil {
                l1 = l1.Next
            } else {
                newL1 := &ListNode{ Val: 0, Next: nil}
                l1.Next = newL1
                l1 = newL1
            }
            if l2.Next != nil {
                l2 = l2.Next
            } else {
                newL2 := &ListNode { Val: 0, Next: nil}
                l2.Next = newL2
                l2 = newL2
            }
        }
    }
    return output.Next
}