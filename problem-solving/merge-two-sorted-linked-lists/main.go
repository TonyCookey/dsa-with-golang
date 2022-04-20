package merge_two_sorted_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1Node *ListNode, list2Node *ListNode) *ListNode {
	if list1Node == nil && list2Node == nil {
		return nil
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := head

	for {
		if list1Node == nil && list2Node == nil {
			break
		}

		if list1Node != nil && list2Node == nil {
			curr.Next = &ListNode{
				Val:  list1Node.Val,
				Next: nil,
			}
			curr = curr.Next
			list1Node = list1Node.Next
		}
		if list1Node == nil && list2Node != nil {
			curr.Next = &ListNode{
				Val:  list2Node.Val,
				Next: nil,
			}
			curr = curr.Next
			list2Node = list2Node.Next
		}
		if list1Node != nil && list2Node != nil {

			if list1Node.Val <= list2Node.Val {

				curr.Next = &ListNode{
					Val:  list1Node.Val,
					Next: nil,
				}
				list1Node = list1Node.Next

			} else {
				curr.Next = &ListNode{
					Val:  list2Node.Val,
					Next: nil,
				}
				list2Node = list2Node.Next
			}
			curr = curr.Next
		}
	}
	return head.Next
}
