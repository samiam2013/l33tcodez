package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nS []*ListNode
	curNode := head
	for curNode != nil {
		nS = append(nS, curNode)
		curNode = curNode.Next
	}
	for i := 0; i < len(nS); i++ {
		if i == len(nS)-n {
			switch {
			case i > 0 && i < len(nS)-1:
				nS[i-1].Next = nS[i+1]
			case i > 0:
				nS[i-1].Next = nil
			case i < len(nS)-1:
				head = nS[i+1]
			default:
				return nil
			}
			break
		}
	}
	return head
}
