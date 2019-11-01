package insertion_sort

func insertionSort(nums []int) {

	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for j >= 0 {
			if temp < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
			j--
		}
		nums[j+1] = temp
	}
}
