package middle_of_the_linked_list

import (
	"fmt"
	"testing"
)

func TestMiddleNode(t *testing.T) {
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
	l3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}
	list := middleNode(l1)
	for list != nil {
		fmt.Print(list.Val, ",")
		list = list.Next
	}
	fmt.Println()
	list = middleNode(l3)
	for list != nil {
		fmt.Print(list.Val, ",")
		list = list.Next
	}
}
