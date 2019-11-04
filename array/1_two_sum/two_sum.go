/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, currentNum := range nums {
		matchNum := target - currentNum
		if value, ok := m[matchNum]; ok {
			return []int{value, index}
		}
		if _, ok := m[currentNum]; !ok {
			m[currentNum] = index
		}
	}
	return []int{}
}
