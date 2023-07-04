package util

import (
	"project/data"
)

func RemoveLinkedListElements(head *data.NodeLeetCode, val int) *data.NodeLeetCode {
	currHead := head
	realhead := head

	if currHead == nil {
		return nil
	}

	if currHead.Val == val {
		return head.Next
	}

	for currHead.Next != nil {
		if currHead.Next.Val == val {
			currHead.Next = currHead.Next.Next
		}
		currHead = currHead.Next
		if currHead == nil {
			break
		}
	}

	return realhead
}
