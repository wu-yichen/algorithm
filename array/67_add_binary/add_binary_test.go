package add_binary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name   string
		num1   string
		num2   string
		expect string
	}{
		{
			name:   "test1",
			num1:   "1011",
			num2:   "1010",
			expect: "10101",
		},
		{
			name:   "test2",
			num1:   "1",
			num2:   "1",
			expect: "10",
		},
		{
			name:   "test3",
			num1:   "101",
			num2:   "1010",
			expect: "1111",
		},
		{
			name:   "test4",
			num1:   "11",
			num2:   "1",
			expect: "100",
		},
	}

	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			actual := addBinary(value.num1, value.num2)
			if actual != value.expect {
				t.Errorf("expect %s, but actual is %s", value.expect, actual)
			}
		})
	}
}
