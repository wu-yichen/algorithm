/*
Given a sorted array nums,
remove the duplicates in-place such that each element appear only once and return the new length.
Do not allocate extra space for another array,
you must do this by modifying the input array in-place with O(1) extra memory.
*/
package remove_duplicates_from_sorted_array

//1,1,2
//0,0,1,1,1,2,2,3,3,4
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
