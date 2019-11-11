package remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name      string
		testCases []int
		expect    int
	}{
		{
			name:      "test1",
			testCases: []int{1, 1, 2},
			expect:    2,
		},
		{
			name:      "test2",
			testCases: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expect:    5,
		},
	}
	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			arr := value.testCases
			actual := removeDuplicates(arr)
			fmt.Println(arr)
			if actual != value.expect {
				t.Error("wrong")
			}
		})
	}
}
