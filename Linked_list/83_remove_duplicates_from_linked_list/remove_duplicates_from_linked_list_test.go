package remove_duplicates_from_linked_list

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	tests := []*ListNode{
		l1, l2,
	}
	for _, value := range tests {
		list := deleteDuplicates(value)
		for list != nil {
			fmt.Print(list.Val, ",")
			list = list.Next
		}
		fmt.Println()
	}

}
