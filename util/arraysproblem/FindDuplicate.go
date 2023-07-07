package arraysproblem

func FindDuplicate(nums []int) int {
	hMap := make(map[int]int)

	for _, v := range nums {
		if _, ok := hMap[v]; ok {
			return v
		}else {
			hMap[v] = 0
		}
	}

	return 0
}