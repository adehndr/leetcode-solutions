package slidingwindow

func MaximumAverageSubarray1(nums []int, k int) float64 {
	return maximumAverageSubarray2(nums, k)
}

// this one failed because time limit exceeded
func maximumAverageSubarray1(nums []int, k int) float64 {
	var result float64
	initValue := 0
	for i := 0; i < k; i++ {
		initValue += nums[i]
	}

	result = float64(initValue) / float64(k)

	for i := 0; i <= len(nums)-k; i++ {
		var tempResult float64
		for j := i; j < i+k; j++ {
			tempResult += float64(nums[j])
		}
		tempResult = tempResult / float64(k)
		if tempResult > result {
			result = tempResult
		}
	}

	return result
}

func abs(input float64) float64 {
	if input < 0 {
		return -1 * input
	}
	return input
}

func maximumAverageSubarray2(nums []int, k int) float64 {
	var result float64
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if i < k {
			result = float64(sum) / float64(k)
		} else {
			sum = sum - nums[i-k]
			if float64(sum)/float64(k) > result {
				result = float64(sum) / float64(k)
			}
		}
	}

	return result
}
