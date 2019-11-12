/*
Given an integer array nums, find the contiguous subarray
(containing at least one number) which has the largest sum and return its sum.
Example:
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/
package maximum_subarray

import "math"

//[-2,1,-3,4,-1,2,1,-5,4], [4,-1,2,1] 6
func maxSubArray(nums []int) int {
	max := int(math.Inf(-1))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 0 {
			sum = 0
			if max < nums[i] {
				max = nums[i]
			}
		} else {
			if max < sum {
				max = sum
			}
		}
	}
	return max
}
