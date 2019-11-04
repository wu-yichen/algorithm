/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You may assume no duplicates in the array.
*/
package search_insert_position

func searchInsert(nums []int, target int) int {
	for key, value := range nums {
		if value >= target {
			return key
		}
	}
	return len(nums)
}
