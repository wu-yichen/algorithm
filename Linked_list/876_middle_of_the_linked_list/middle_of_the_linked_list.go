package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//func middleNode(head *ListNode) *ListNode {
//	set := make(map[int]*ListNode)
//	count := 1
//	for head != nil {
//		set[count] = head
//		head = head.Next
//		count++
//	}
//	num := count - 1
//	return set[num/2+1]
//}
