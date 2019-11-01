/*
Given an array nums and a value val, remove all instances of that value in-place and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
The order of elements can be changed. It doesn't matter what you leave beyond the new length.
*/
package remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	j := len(nums) - 1
	for key, value := range nums {
		if value == val {
			for j >= 0 && nums[j] == val {
				if j == key {
					return len(nums[:key])
				}
				j--
			}
			nums[key], nums[j] = nums[j], nums[key]
		}
	}
	return len(nums)
}
