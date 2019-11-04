package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		expect []int
	}{
		{
			name:   "test1",
			input:  []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		{
			name:   "test2",
			input:  []int{2, 7, 11, 15},
			target: 18,
			expect: []int{1, 2},
		},
		{
			name:   "test3",
			input:  []int{2, 7, 11, 15},
			target: 100,
			expect: []int{},
		},
		{
			name:   "test4",
			input:  []int{},
			target: 100,
			expect: []int{},
		},
		{
			name:   "test5",
			input:  nil,
			target: 100,
			expect: []int{},
		},
	}
	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			actual := twoSum(value.input, value.target)
			expect := value.expect
			for key, value := range expect {
				if value != actual[key] {
					t.Errorf("function twoSum = %v, actual %v ", expect, actual)
				}
			}
		})
	}
}
