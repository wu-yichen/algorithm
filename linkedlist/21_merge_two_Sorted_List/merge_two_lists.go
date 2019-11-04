/*
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.
Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
package merge_two_Sorted_List

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil{
//		return l2
//	}
//	if l2 == nil{
//		return l1
//	}
//	if l1.Val < l2.Val{
//		l1.Next = mergeTwoLists(l1.Next,l2)
//		return l1
//	}
//	l2.Next = mergeTwoLists(l1,l2.Next)
//	return l2
//}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tail.Next = l2
			break
		}
		if l2 == nil {
			tail.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1, tail = l1.Next, tail.Next
		} else {
			tail.Next = l2
			l2, tail = l2.Next, tail.Next
		}
	}
	return dummy.Next
}
