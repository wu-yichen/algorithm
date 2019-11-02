package delete_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) { //nolint
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
