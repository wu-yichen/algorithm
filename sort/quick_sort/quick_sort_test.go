package quick_sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	target := []int{3, 4, 23, 56, 456, 67, 87634, 234545, 65, 42, 1, 34, 4, 56, 7}
	quickSort(target, 0, len(target)-1)
	expect := []int{1, 3, 4, 4, 7, 23, 34, 42, 56, 56, 65, 67, 456, 87634, 234545}

	if !reflect.DeepEqual(target, expect) {
		t.Errorf("test failed,\n expect \n %v \n but actual is \n %v", expect, target)
	}
}
