package palindrome_linked_list

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	l3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	l4 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l5 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	tests := []struct {
		name   string
		n      *ListNode
		expect bool
	}{
		{
			name:   "test1",
			n:      l1,
			expect: true,
		},
		{
			name:   "test2",
			n:      l2,
			expect: true,
		},
		{
			name:   "test3",
			n:      l3,
			expect: false,
		},
		{
			name:   "test4",
			n:      l4,
			expect: false,
		},
		{
			name:   "test5",
			n:      l5,
			expect: false,
		},
	}
	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			palindrome := isPalindrome(value.n)
			if palindrome != value.expect {
				t.Errorf("%s is wrong, expect %v, but actually %v", value.name, value.expect, palindrome)
			}
		})
	}
}
