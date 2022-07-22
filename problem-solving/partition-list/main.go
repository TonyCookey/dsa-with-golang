package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var hasSeen bool
	dummyHead := &ListNode{
		Next: &ListNode{},
	}
	dummyHead.Next.Next = head
	lastLess := dummyHead.Next
	prev := dummyHead.Next
	for head != nil {
		if head.Val < x {
			if hasSeen {
				temp := lastLess.Next
				lastLess.Next = head
				next := head.Next
				head.Next = temp
				prev.Next = next
				lastLess = head
				head = next
				continue
			}
			lastLess = head
		}
		if head.Val >= x {
			hasSeen = true
		}
		prev = head
		head = head.Next
	}
	return dummyHead.Next.Next
}
