package reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	nodeStore := make(map[int]*ListNode)
	dummyHead := head
	i := 1
	for dummyHead != nil {
		nodeStore[i] = dummyHead
		dummyHead = dummyHead.Next
		i++
	}

	for left < right {
		temp := nodeStore[left].Val
		nodeStore[left].Val = nodeStore[right].Val
		nodeStore[right].Val = temp
		left++
		right--
	}
	return head
}
