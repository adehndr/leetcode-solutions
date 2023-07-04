package util

import "project/data"

func IsPalindromLinkedList(head *data.NodeLeetCode) bool {
	currHead := head
	aList := []int{}
	for currHead != nil {
		aList = append(aList, currHead.Val)
		currHead = currHead.Next
	}
	for i := 0; i < len(aList)/2; i++ {
		if aList[i] != aList[len(aList)-1-i] {
			return false
		}
	}

	return true
}
