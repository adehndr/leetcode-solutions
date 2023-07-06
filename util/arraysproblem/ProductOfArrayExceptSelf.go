package arraysproblem

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	hMap := make(map[int]int)

	for k, v := range nums {
		hMap[k] = v
	}

	for i := range result {
		tempRes := 1
		for k, v := range hMap {
			if k != i {
				tempRes = tempRes * v
			}
		}
		result[i] = tempRes
	}

	return result
}

func ProductExceptSelf2(nums []int) []int {
	result := make([]int, len(nums))
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

func ProductExceptSelfNeetCode(nums []int) []int {
	prefix := []int{}
	postfix := []int{}
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i <= 0 {
			prefix = append(prefix, nums[i])
		} else {
			prefix = append(prefix, nums[i]*prefix[i-1])
		}
		if len(postfix) == 0 {
			postfix = append([]int{nums[len(nums)-i-1]}, postfix...)
		} else {
			postfix = append([]int{nums[len(nums)-i-1] * postfix[len(postfix)-i]}, postfix...)
		}
	}

	for i := 0; i < len(nums); i++ {
		if i <= 0 || i+1 >= len(nums) {
			if i <= 0 {
				result[i] = postfix[i+1]
			} else {
				result[i] = prefix[i-1]
			}
		} else {
			result[i] = prefix[i-1] * postfix[i+1]
		}
	}

	return result
}

func ProductExceptSelfNeetCode2(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix = prefix * nums[i]
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * postfix
		postfix = postfix * nums[i]
	}
	return res
}
