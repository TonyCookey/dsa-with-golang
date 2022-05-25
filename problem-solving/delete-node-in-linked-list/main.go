package delete_node_in_linked_list

type ListNode struct {
	Next *ListNode
	Val  int
}

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}
