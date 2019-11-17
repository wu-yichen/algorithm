package find_duplication

// 2,3,1,0,2,5  - 2
func findDuplication(input []int) int {
	if len(input) == 0 {
		return -1
	}
	for key, value := range input {
		if value != key {
			if value == input[value] {
				return value
			}
			input[key], input[input[key]] = input[input[key]], input[key]
		}
	}
	return -1
}
