package reverse_linked_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	}

	list := reverseList(l1)
	for list != nil {
		fmt.Print(list.Val, ",")
		list = list.Next
	}
}
