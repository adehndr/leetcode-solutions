package arraysproblem

func findDuplicates(nums []int) []int {
	hMap := make(map[int]int)
	result := []int{}
	for _, ele := range nums {
		if _, ok := hMap[ele]; ok {
			result = append(result, ele)
		}else {
			hMap[ele] = 0
		}
	}

	return result
}
