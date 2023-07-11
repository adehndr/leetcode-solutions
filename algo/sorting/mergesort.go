package sorting

func MergeSort(nums []int) {

	if len(nums) <= 1 {
		return
	}

	mid := len(nums) / 2
	tempLeft := []int{}
	tempRight := []int{}
	tempLeft = append(tempLeft, nums[:mid]...)
	tempRight = append(tempRight, nums[mid:]...)

	MergeSort(tempLeft)
	MergeSort(tempRight)
	merge(nums, tempLeft, tempRight)
}

func merge(arr, left, right []int) {
	l := 0
	r := 0
	k := 0

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			arr[k] = right[r]
			r += 1
			k += 1
		} else {
			arr[k] = left[l]
			l += 1
			k += 1
		}
	}

	for l < len(left) {
		arr[k] = left[l]
		l += 1
		k += 1
	}

	for r < len(right) {
		arr[k] = right[r]
		r += 1
		k += 1
	}

}
