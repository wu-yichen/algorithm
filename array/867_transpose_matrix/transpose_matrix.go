package transpose_matrix

func transpose(A [][]int) [][]int {
	result := make([][]int, len(A[0]))
	for index := 0; index < len(A[0]); index++ {
		result[index] = make([]int, len(A))
		for i := range A {
			result[index][i] = A[i][index]
		}
	}
	return result
}
