package arraysproblem

func ProductExceptSelf(nums []int) []int {
	result := make([]int,len(nums))
	hMap := make(map[int]int)

	for k , v := range nums {
		hMap[k] = v
	}

	for i := range result {
		tempRes := 1
		for k , v := range hMap {
			if k != i {
				tempRes = tempRes * v
			}
		}
		result[i] = tempRes
	}

	return result
}

func ProductExceptSelf2(nums []int) []int {
	result := make([]int,len(nums))
	for k := range nums {
		result[k] = recAccessSlice(nums[:k]) * recAccessSlice(nums[k+1:])
	}
	return result
}

func recAccessSlice(slice []int) int {
	if len(slice) == 0 {
		return 1
	}
	return slice[0] * recAccessSlice(slice[1:])
}
