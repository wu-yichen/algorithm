package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		val    int
		expect int
	}{
		{
			name:   "test1",
			input:  []int{3, 2, 2, 3},
			val:    3,
			expect: 2,
		},
		{
			name:   "test2",
			input:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			expect: 5,
		},
		{
			name:   "test3",
			input:  []int{0},
			val:    0,
			expect: 0,
		},
		{
			name:   "test4",
			input:  []int{0},
			val:    2,
			expect: 1,
		},
		{
			name:   "test5",
			input:  []int{0, 1},
			val:    0,
			expect: 1,
		},
		{
			name:   "test6",
			input:  []int{},
			val:    0,
			expect: 0,
		},
		{
			name:   "test7",
			input:  []int{3, 3},
			val:    3,
			expect: 0,
		},
		{
			name:   "test8",
			input:  []int{3, 3, 3},
			val:    5,
			expect: 3,
		},
	}
	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			actual := removeElement(value.input, value.val)
			expect := value.expect
			if actual != expect {
				t.Errorf("expect %v, but actual %v", expect, actual)
			}
		})
	}
}
