package util

func Convert1DTo2D(original []int, m int, n int) [][]int {
	result := [][]int{}

	if len(original) != m*n {
		return result
	}

	k := 0
	for i := 0; i < m; i++ {
		result = append(result, []int{})
		for j := 0; j < n; j++ {
			result[i] = append(result[i], original[k])
			k += 1
		}
	}

	return result
}
