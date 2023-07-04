package util

import (
	"strconv"
)

func CountBits(n int) []int {
	var result []int

	// initiate
	for i := 0; i <= n; i++ {
		result = append(result, 0)
	}

	var resultBit []string
	i := 0

	for i <= n {
		resultBit = append(resultBit, strconv.FormatInt(int64(i), 2))
		i++
	}
	for k, v := range resultBit {
		for i := 0; i < len(v); i++ {
			if v[i] == 49 {
				result[k] += 1
			}
		}
	}

	return result
}
