package search

func BinarySearch(input []int, target int) int {
	return binarySearchOptimized(input, target)
}

func binarySearchOptimized(input []int, target int) int {
	low := 0
	high := len(input) - 1
	result := -1
	mid := low + ((high - low) / 2)
	for low <= high {
		if target > input[mid] {
			low = mid + 1
		} else if target < input[mid] {
			high = mid - 1
		} else {
			result = mid
			return result
		}
		mid = low + ((high - low) / 2)
	}

	return result
}
