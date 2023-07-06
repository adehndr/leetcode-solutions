package binarytreeproblem

import (
	"fmt"
)

func BinaryTreeLevelOrderTraversal(root *treeNode) [][]int {
	aQueue := queueTreeNode{}
	aQueue.enque(root)
	finalRes := [][]int{}

	for !aQueue.isEmpty() {
		lengthQ := aQueue.length()
		tempRes := []int{}
		fmt.Println(lengthQ)
		for i := 0; i < lengthQ; i++ {
			tempDeq := aQueue.dequeue()
			fmt.Println(tempDeq)
			if tempDeq.Left != nil {
				aQueue.enque(tempDeq.Left)
			}
			if tempDeq.Right != nil {
				aQueue.enque(tempDeq.Right)
			}
			tempRes = append(tempRes, tempDeq.Val)
		}
		finalRes = append(finalRes, tempRes)
	}

	return finalRes
}

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

type queueTreeNode struct {
	data []treeNode
}

func (q *queueTreeNode) enque(input *treeNode) {
	if input != nil {
		(*q).data = append((*q).data, *input)
	}
}

func (q *queueTreeNode) dequeue() *treeNode {
	if q.isEmpty() {
		return nil
	}

	temp := (*q).data[0]
	(*q).data = (*q).data[1:]
	return &temp
}

func (q *queueTreeNode) isEmpty() bool {
	return len((*q).data) == 0
}

func (q *queueTreeNode) length() int {
	return len((*q).data)
}
