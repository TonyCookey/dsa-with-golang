package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		return slow.Next
	} else {
		return slow
	}
}

func middleNode2(head *ListNode) *ListNode {
	s := []*ListNode{}
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	return s[len(s)/2]
}
