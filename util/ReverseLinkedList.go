package util

import (
	"project/data"
)

func ReverseList(head *data.NodeLeetCode) *data.NodeLeetCode {
	if head == nil {
		return nil
	}

	hMap := make(map[int]*data.NodeLeetCode)
	aList := []int{}
	currHead := head
	count := 0
	for currHead != nil {
		hMap[count] = currHead
		aList = append(aList, count)
		currHead = currHead.Next
		count += 1
	}

	mergeSortDesc(aList)
	for i := 0; i < len(aList)-1; i++ {
		tmpNode := hMap[aList[i]]
		tmpNode.Next = hMap[aList[i+1]]
	}
	hMap[0].Next = nil
	return hMap[len(aList)-1]
}

func mergeSortDesc(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	tmpLeft := []int{}
	tmpRight := []int{}

	tmpLeft = append(tmpLeft, arr[:mid]...)
	tmpRight = append(tmpRight, arr[mid:]...)

	mergeSortDesc(tmpLeft)
	mergeSortDesc(tmpRight)

	sortDesc(arr, tmpLeft, tmpRight)

}

func sortDesc(arr, left, right []int) {
	l := 0
	i := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
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
