package binarysearch

import "fmt"

func FindPeakElement(nums []int) int {
	return findPeakElementOptimized(nums)
}

func findPeakElement1(nums []int) int {
	var result int = 0
	if len(nums) == 0 {
		return result
	} else if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			result = nums[0]
		} else {
			result = nums[1]
		}
		return result
	}

	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	for left <= right {
		fmt.Println(nums[left], nums[right], nums[mid])
		if mid-1 >= 0 && nums[mid-1] > nums[mid] {
			right = mid - 1
		} else if mid+1 <= len(nums)-1 && nums[mid+1] > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
		mid = left + (right-left)/2
	}
	return result
}

func findPeakElementOptimized(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	for left <= right {
		if mid > 0 && nums[mid-1] > nums[mid] {
			right = mid - 1
		} else if mid < len(nums)-1 && nums[mid+1] > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
		mid = left + (right-left)/2
	}
	return -1
}
