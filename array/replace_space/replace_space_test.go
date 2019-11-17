package replace_space

import "testing"

func TestReplaceSpace(t *testing.T) {
	test := "A B"
	actual := replaceSpace(test)
	expected := "A%20B"
	if actual != expected {
		t.Errorf("wrong ! expectd %v, but actual %v", expected, actual)
	}
}
