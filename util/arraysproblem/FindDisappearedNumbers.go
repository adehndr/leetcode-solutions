package arraysproblem

func findDisappearedNumbers(nums []int) []int {
	tmpMap := make(map[int]int)
	result := []int{}

	for i := 1; i <= len(nums); i++ {
		tmpMap[i] = 0
	}

	for _, v := range nums {
		if _, ok := tmpMap[v]; ok {
			tmpMap[v] = 1
		}
	}

	for k, v := range tmpMap {
		if v == 0 {
			result = append(result, k)
		}
	}

	return result
}

func FindDisappearedNumbers2(nums []int) []int {
	hMap := make(map[int]int)
	for _, v := range nums {
		hMap[v] = 1
	}
	result := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := hMap[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}

func FindDisappearedNumbersNeetCode(nums []int) []int {
	result := []int{}
	for _, v := range nums {
		i := abs(v) - 1
		nums[i] = -1 * abs(nums[i])
	}

	for k, v := range nums {
		if v > 0 {
			result = append(result, k+1)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
