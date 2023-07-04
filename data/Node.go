package data

import "fmt"

type Node struct {
	Tail *Node
	Data string
}

type NodeLeetCode struct {
	Val  int
	Next *NodeLeetCode
}

func (nLC *NodeLeetCode) TraverseNodeLeetCode() {
	currHead := nLC
	for currHead != nil {
		fmt.Printf("%d ->",currHead.Val)
		currHead = currHead.Next
	}
	fmt.Println()
}

func (lN *NodeLeetCode) addNodeLeetCode(node *NodeLeetCode) {
	currNode := lN
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = node
}

func InitNodeLeetCode(nums []int) *NodeLeetCode {
	if len(nums) <= 0 {
		return nil
	}
	initNode := &NodeLeetCode{
		Val: nums[0],
	}
	
	currNode := initNode
	for i := 1; i < len(nums); i++ {
		tmpNode := &NodeLeetCode{
			Val: nums[i],
		}
		currNode.addNodeLeetCode(tmpNode)
		currNode = currNode.Next
	}
	return initNode
}