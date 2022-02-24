package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// sort a singly linked list using merge sort
func sortList(head *ListNode) *ListNode {
	// using merge sort

	if head == nil || head.Next == nil {
		return head
	}
	// split the linked list into two halves
	left := head
	right := getMiddleNode(head)
	temp := right.Next
	right.Next = nil
	right = temp

	left = sortList(left)
	right = sortList(right)

	return mergeList(left, right)

}

func getMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(left, right *ListNode) *ListNode {
	var dummy = &ListNode{2, nil}

	tail := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}
		tail = tail.Next
	}
	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}

	return dummy.Next
}
