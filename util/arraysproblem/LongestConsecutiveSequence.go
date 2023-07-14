package arraysproblem

import "fmt"

func LongestConsecutiveSequence(nums []int) int {
	return longestConsecutiveSequence2(nums)
}

func longestConsecutiveSequence(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var result int = 1
	var count int = 1
	mergeSort(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}

		if nums[i+1] == nums[i]+1 {
			count += 1
			if count >= result {
				result = count
			}
		} else {
			count = 1
		}
	}

	return result
}

func mergeSort(nums []int) {

	if len(nums) <= 1 {
		return
	}

	mid := len(nums) / 2

	tempLeft := []int{}
	tempRight := []int{}

	tempLeft = append(tempLeft, nums[:mid]...)
	tempRight = append(tempRight, nums[mid:]...)

	mergeSort(tempLeft)
	mergeSort(tempRight)
	merge(nums, tempLeft, tempRight)
}

func merge(arr, left, right []int) {
	l := 0
	r := 0
	i := 0

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			arr[i] = right[r]
			r += 1
			i += 1
		} else {
			arr[i] = left[l]
			l += 1
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

func longestConsecutiveSequence2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var longest int
	hMap := make(map[int]bool)

	for _, v := range nums {
		hMap[v] = true
	}

	for _, v := range nums {
		if !hMap[v-1] {
			currentCount := 1
			currentVal := v
			for hMap[currentVal+1] {
				currentCount += 1
				currentVal += 1
			}

			if currentCount >= longest {
				longest = currentCount
			}
		}
	}

	return longest
}
