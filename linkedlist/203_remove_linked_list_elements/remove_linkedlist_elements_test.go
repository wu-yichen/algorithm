package remove_linked_list_elements

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	l1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	list := removeElements(l1, 5)
	for list != nil {
		fmt.Print(list.Val, ",")
		list = list.Next
	}
}
