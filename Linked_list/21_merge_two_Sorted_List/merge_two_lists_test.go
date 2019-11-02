package merge_two_Sorted_List

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list := mergeTwoLists(l1, l2)
	for list != nil {
		fmt.Print(list.Val, ",")
		list = list.Next
	}
}
