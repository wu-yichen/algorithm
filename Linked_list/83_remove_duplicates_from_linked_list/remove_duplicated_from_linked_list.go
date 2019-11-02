package remove_duplicates_from_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	dummy := current
	for current != nil && current.Next != nil {
		nxt := current.Next
		if nxt.Val == current.Val {
			current.Next = nxt.Next
			continue
		}
		current = nxt
	}
	return dummy
}
