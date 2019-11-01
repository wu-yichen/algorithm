package bubble_sort

func bubbleSort(nums []int) {
	i := len(nums)
	for i > 0 {
		swapped := false
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		i--
	}
}
