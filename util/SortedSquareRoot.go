package util

func SortedSquareRoot(nums []int) []int {
	return sortedSquareRoot2(nums)
}

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	tmpLeft := []int{}
	tmpRight := []int{}

	tmpLeft = append(tmpLeft, arr[:mid]...)
	tmpRight = append(tmpRight, arr[mid:]...)

	mergeSort(tmpLeft)
	mergeSort(tmpRight)

	sort(arr, tmpLeft, tmpRight)

}

func sort(arr, left, right []int) {
	l := 0
	i := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			arr[i] = left[l]
			l += 1
			i += 1
		} else {
			arr[i] = right[r]
			r += 1
			i += 1
		}
	}
	for l < len(left) {
		arr[i] = left[l]
		l += 1
		i += 1
	}

	for r < len(right) {
		arr[i] = right[r]
		r += 1
		i += 1
	}
}

func sortedSquareRoot(nums []int) []int {
	result := []int{}

	for _, v := range nums {
		result = append(result, v*v)
	}
	mergeSort(result)

	return result
}

func sortedSquareRoot2(nums []int) []int {
	result := make([]int, len(nums))
	end := len(result) - 1
	l := 0
	r := len(nums) - 1

	for l <= r {
		if abs(nums[l]) < abs(nums[r]) {
			result[end] = nums[r] * nums[r]
			r -= 1
			} else {
			result[end] = nums[l] * nums[l]
			l += 1
		}
		end -= 1
	}

	return result
}

func abs(input int) int {
	if input < 0 {
		return -1 * input
	}
	return input
}
