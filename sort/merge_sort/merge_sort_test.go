package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{1, 3, 45, 235, 65, 234, 43, 21, 2}
	sort := mergeSort(nums)
	expect := []int{1, 2, 3, 21, 43, 45, 65, 234, 235}
	if !reflect.DeepEqual(expect, sort) {
		t.Error("not equal")
	}
}
