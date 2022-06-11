package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var count int
	curr := head

	for curr != nil {
		count++
		curr = curr.Next
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	curr = dummyHead
	for i := 0; i <= count-n; i++ {
		if i == count-n {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return dummyHead.Next
}
