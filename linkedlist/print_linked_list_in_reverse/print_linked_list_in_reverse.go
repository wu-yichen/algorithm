package print_linked_list_in_reverse

type ListNode struct {
	Val  int
	Next *ListNode
}

//1->2->3
func reverse(listNode *ListNode) *ListNode {
	if listNode == nil {
		return nil
	}
	head := &ListNode{}
	head.Next = nil

	for listNode != nil {
		temp := listNode.Next
		listNode.Next = head.Next
		head.Next = listNode
		listNode = temp
	}

	return head.Next
}
