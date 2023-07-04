package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (lN *ListNode) AddNode(node *ListNode) {
	currNode := lN
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = node
}

func InitNode(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	initNode := &ListNode{
		Val: nums[0],
	}
	
	currNode := initNode
	for i := 1; i < len(nums); i++ {
		tmpNode := &ListNode{
			Val: nums[i],
		}
		currNode.AddNode(tmpNode)
		currNode = currNode.Next
	}
	return initNode
}

func (lN *ListNode) Traverse() {
	currNode := lN
	fmt.Printf("%v %d\n", currNode, currNode.Val)
	for currNode.Next != nil {
		fmt.Printf("%v %d\n", currNode.Next, currNode.Next.Val)
		currNode = currNode.Next
	}
}

func HasCycler(head *ListNode) bool {
	hMap := map[*ListNode]int{}
	currNode := head
	for currNode.Next != nil {
		if _, ok := hMap[currNode]; ok {
			return true
		} else {
			hMap[currNode] = 1
		}
		currNode = currNode.Next
	}
	return false
}

func (lN *ListNode) GetNNode(n int) *ListNode{
	currNode := lN
	for i := 0; i < n; i++ {
		currNode = currNode.Next
	}
	return currNode
}
