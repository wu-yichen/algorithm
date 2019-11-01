package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		expect int
	}{
		{
			name:   "test1",
			input:  []int{1, 3, 5, 6},
			target: 5,
			expect: 2,
		},
		{
			name:   "test2",
			input:  []int{1, 3, 5, 6},
			target: 2,
			expect: 1,
		},
		{
			name:   "test3",
			input:  []int{1, 3, 5, 6},
			target: 7,
			expect: 4,
		},
		{
			name:   "test4",
			input:  []int{},
			target: 7,
			expect: 0,
		},
		{
			name:   "test5",
			input:  []int{7},
			target: 7,
			expect: 0,
		},
	}

	for key, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			actual := searchInsert(value.input, value.target)
			expect := value.expect
			if actual != expect {
				t.Errorf("failed on %vth test, expect %v, but actual is %v", key+1, expect, actual)
			}
		})
	}
}
