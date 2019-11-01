package quick_sort

func quickSort(target []int, lowest int, highest int) {
	if lowest < highest {
		// get partition
		p := partition(target, lowest, highest)

		//recursive divide and conquer
		quickSort(target, lowest, p-1)
		quickSort(target, p+1, highest)
	}
}

func partition(target []int, lowest int, highest int) int {
	pivot := target[highest]
	i := lowest - 1
	for j := lowest; j < highest; j++ {
		if target[j] <= pivot {
			i++
			target[i], target[j] = target[j], target[i]
		}
	}
	target[i+1], target[highest] = target[highest], target[i+1]
	return i + 1
}
