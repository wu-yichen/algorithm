/*
Given a linked list, determine if it has a cycle in it.
To represent a cycle in the given linked list,
we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to.
If pos is -1, then there is no cycle in the linked list.
Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.
*/
package Linked_List_Cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

//func hasCycle(head *ListNode) bool {
//	set := make(map[*ListNode]bool)
//	for head != nil {
//		if ok := set[head]; ok{
//			return true
//		}
//		set[head] = true
//		head = head.Next
//	}
//	return false
//}
func hasCycle(head *ListNode) bool { //nolint
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next.Next
	slow := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
