package remove_duplicates_from_sorted_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	curr := head.Next

	for curr != nil {
		if curr.Val != prev.Val {
			prev.Next = curr
			prev = curr
		}
		curr = curr.Next
	}
	prev.Next = nil

	return head
}
