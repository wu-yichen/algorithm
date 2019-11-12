package maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		testCase []int
		expect   int
	}{
		{
			name:     "test1",
			testCase: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect:   6,
		},
		{
			name:     "test2",
			testCase: []int{-1},
			expect:   -1,
		},
	}

	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			actual := maxSubArray(value.testCase)
			if actual != value.expect {
				t.Errorf("wrong: the actual value is %v", actual)
			}
		})
	}
}
