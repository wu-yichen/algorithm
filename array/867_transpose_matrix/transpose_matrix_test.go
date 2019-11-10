package transpose_matrix

import (
	"fmt"
	"testing"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		name     string
		testCase [][]int
		expect   [][]int
	}{
		{
			name: "test1",
			testCase: [][]int{
				{2, 4, -1},
				{-10, 5, 11},
				{18, -7, 6},
			},
			expect: [][]int{
				{2, -10, -18},
				{4, 5, -7},
				{-1, 11, 6},
			},
		},
		{
			name: "test2",
			testCase: [][]int{
				{2, 4, -1, 9},
				{-10, 5, 11, 10},
				{18, -7, 6, 11},
				{1, 2, 3, 4},
			},
			expect: [][]int{
				{2, -10, 18, 1},
				{4, 5, -7, 2},
				{-1, 11, 6, 3},
				{9, 10, 11, 4},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := transpose(test.testCase)
			for _, value := range actual {
				for _, val := range value {
					fmt.Print(val, "  ")
				}
				fmt.Println()
			}
		})
	}
}
