package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var store []int
	for head != nil {
		store = append(store, head.Val)
		head = head.Next
	}
	i, j := 0, len(store)-1
	for i < j {
		if store[i] != store[j] {
			return false
		}
		i++
		j--
	}
	return true
}
