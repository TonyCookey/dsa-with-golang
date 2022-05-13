package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	store := make(map[*ListNode]bool)

	for head != nil {
		if store[head] {
			return true
		} else {
			store[head] = true
			head = head.Next
		}
	}
	return false
}
func hasCycleConstantSpace(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
