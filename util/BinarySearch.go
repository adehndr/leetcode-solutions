package util

import "fmt"

func BinarySearch(nums []int, target int) int {
	result := -1

	left := 0
	right := len(nums) - 1
	// -1,0,3,5,9,12
	for left <= right {
		mid := left + ((right - left) / 2)
		fmt.Println(mid)
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			result = mid
			return result
		}
	}
	return result
}
