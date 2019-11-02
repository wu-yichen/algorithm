package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	current := new(ListNode)
	ptr := current
	current.Next = head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return ptr.Next
}

//func removeElements(head *ListNode, val int) *ListNode {
//	res := &ListNode{}
//	tail := res
//	for head != nil {
//		if head.Val != val {
//			tail.Next = head
//			tail = tail.Next
//		}
//		head = head.Next
//	}
//	tail.Next = nil
//	return res.Next
//}
