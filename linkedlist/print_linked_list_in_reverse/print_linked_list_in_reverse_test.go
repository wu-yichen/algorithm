package print_linked_list_in_reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		testCase *ListNode
		expected *ListNode
	}{
		{
			name: "test1",
			testCase: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 3,
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

	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			actual := reverse(value.testCase)
			expect := value.expected
			for actual != nil {
				if actual.Val != expect.Val {
					t.Errorf("wrong, expect %v, but %v", expect.Val, actual.Val)
				}
				expect = expect.Next
				actual = actual.Next
			}
		})
	}
}
