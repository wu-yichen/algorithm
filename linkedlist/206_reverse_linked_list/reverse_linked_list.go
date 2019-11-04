/*
Reverse a singly linked list.
Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/
package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	current := head
	var previous *ListNode = nil
	for current != nil {
		nxt := current.Next
		current.Next = previous
		previous = current
		current = nxt
	}
	return previous
}
