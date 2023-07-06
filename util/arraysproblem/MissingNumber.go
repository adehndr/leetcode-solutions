package arraysproblem

func MissingNumber(nums []int) int {
	hMap := make(map[int]int)

	for _, v := range nums {
		hMap[v] = 0
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := hMap[i]; !ok {
			return i
		}
	}

	return -1
}
