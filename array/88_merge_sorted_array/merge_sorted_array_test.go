package merge_sorted_array

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge(arr1, 3, arr2, 3)
	for _, value := range arr1 {
		fmt.Println(value)
	}
}
