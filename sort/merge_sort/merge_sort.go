package merge_sort

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	m := len(nums) / 2
	return merge(mergeSort(nums[:m]), mergeSort(nums[m:]))
}
func merge(subArray1 []int, subArray2 []int) []int {
	result := make([]int, len(subArray1)+len(subArray2))
	i, j, k := 0, 0, 0

	for k < len(result) {
		if i < len(subArray1) && j >= len(subArray2) {
			result[k] = subArray1[i]
			i++
		} else if i >= len(subArray1) && j < len(subArray2) {
			result[k] = subArray2[j]
			j++
		} else if subArray1[i] < subArray2[j] {
			result[k] = subArray1[i]
			i++
		} else {
			result[k] = subArray2[j]
			j++
		}
		k++
	}
	return result
}
