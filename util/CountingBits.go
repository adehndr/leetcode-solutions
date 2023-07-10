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

func CountBits2(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		tempRes := strconv.FormatInt(int64(i), 2)
		count := 0
		for _, v := range tempRes {
			if string(v) == "1" {
				count += 1
			}
		}
		result[i] = count
	}

	return result
}

func CountBits3(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		curr := i
		count := 0
		for curr > 0 {
			if curr % 2 == 1 {
				count += 1
			}
			curr = curr / 2
		}
		result[i] = count
	}

	return result
}
