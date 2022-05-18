package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		if curr.Val == val && prev == nil {
			head = curr.Next
			curr = head
			continue
		} else {
			if curr.Val == val && prev != nil {
				prev.Next = curr.Next
				curr = curr.Next
				continue
			}
			prev = curr
			curr = curr.Next
		}
	}
	return head
}
