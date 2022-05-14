package interval_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	seenNodeMap := make(map[*ListNode]bool)

	for headA != nil || headB != nil {
		if seenNodeMap[headA] {
			return headA
		}
		if headA != nil {
			seenNodeMap[headA] = true
			headA = headA.Next
		}
		if seenNodeMap[headB] {
			return headB
		}
		if headB != nil {
			seenNodeMap[headB] = true
			headB = headB.Next
		}
	}
	return nil
}
