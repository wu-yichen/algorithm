package find_duplication

import "testing"

//// 2,3,1,0,2,5  - 2
func TestFindDuplication(t *testing.T) {
	test := []int{2, 3, 1, 0, 2, 5}
	actual := findDuplication(test)
	if actual != 2 {
		t.Errorf("wrong! expect %v, but actual %v", 2, actual)
	}
}
