package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}
	var carry int

	head := &ListNode{
		Val: 0,
	}
	dummyhead := head

	for l1 != nil || l2 != nil {
		var add int

		if l1 != nil {
			add += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			add += l2.Val
			l2 = l2.Next
		}
		if carry == 1 {
			add += carry
			carry = 0
		}
		if add > 9 {
			carry = add / 10
			add = add % 10
		}
		head.Next = &ListNode{
			Val: add,
		}
		head = head.Next
	}
	if carry == 1 {
		head.Next = &ListNode{
			Val: carry,
		}
	}
	return dummyhead.Next
}
