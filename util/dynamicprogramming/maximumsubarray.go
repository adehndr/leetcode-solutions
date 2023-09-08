package dynamicprogramming

func MaximumSubarray(nums []int) int {
	return maximumSubarrayQuadratic(nums)
}

func maximumSubarrayQuadratic(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		var tmp int
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			result = findMax(result, tmp)
		}
	}
	return result
}

func maximumSubarrayOptimum(nums []int) int {
	maxSub := nums[0]
	var result int

	for _, v := range nums {
		result += v
		maxSub = findMax(maxSub, result)
		if maxSub < 0 {
			result = 0
		}
	}

	return result
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
