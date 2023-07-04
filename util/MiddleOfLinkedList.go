package util

import "project/data"

func MiddleNode(head *data.NodeLeetCode) *data.NodeLeetCode {
	count := 0
	currHead := head
	var result *data.NodeLeetCode = nil

	for currHead != nil {
		count += 1
		currHead = currHead.Next
	}

	mid := count / 2
	i := 0
	currHead = head
	result = currHead
	for currHead != nil {
		if i == mid {
			return result
		}
		currHead = currHead.Next
		result = currHead
		i += 1
	}
	return result
}