package arraysproblem

func missingNumber(nums []int) int {
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

func missingNumber2(nums []int) int {
	result := (len(nums) * (len(nums) + 1) ) / 2
	for _, v := range nums {
		result -= v
	}
	return result
}

func MissingNumber(nums []int) int {
	return missingNumber2(nums)
}
