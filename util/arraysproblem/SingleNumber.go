package arraysproblem

func singleNumber(nums []int) int {
	tmpMap := make(map[int]int)

	for _, v := range nums {
		if _, ok := tmpMap[v]; ok {
			tmpMap[v] = 2
		} else {
			tmpMap[v] = 1
		}
	}

	for k, v := range tmpMap {
		if v == 1 {
			return k
		}
	}

	return -1

}

func singleNumberBitManipulation(nums []int) int {
	var result int
	for _, v := range nums {
		result ^= v
	}

	return result
}

func SingleNumberBitManipulation(input []int) int {
	return singleNumberBitManipulation(input)
}
