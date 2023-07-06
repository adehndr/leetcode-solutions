package data

import "fmt"

type NodeBinarySearchTree struct {
	Data  int
	left  *NodeBinarySearchTree
	right *NodeBinarySearchTree
}

type NodeBinarySearchTreeLeetCode struct {
	Val   int
	Left  *NodeBinarySearchTreeLeetCode
	Right *NodeBinarySearchTreeLeetCode
}

type BinarySearchTree struct {
	node *NodeBinarySearchTree
}

func (bST *BinarySearchTree) Insert(newNode *NodeBinarySearchTree) {
	bST.node = bST.insertHelper(bST.node, newNode)
}

func (bST *BinarySearchTree) insertHelper(
	rootNode *NodeBinarySearchTree,
	newNode *NodeBinarySearchTree,
) *NodeBinarySearchTree {
	tempData := newNode.Data

	if rootNode == nil {
		rootNode = newNode
		return rootNode
	} else if rootNode.Data > tempData {
		rootNode.left = bST.insertHelper(rootNode.left, newNode)
	} else {
		rootNode.right = bST.insertHelper(rootNode.right, newNode)
	}

	return rootNode
}

func (bST *BinarySearchTree) Search(num int) bool {
	return bST.searchHelper(bST.node, num)
}

func (bST *BinarySearchTree) searchHelper(rootNode *NodeBinarySearchTree, num int) bool {
	if rootNode == nil {
		return false
	} else if rootNode.Data > num {
		return bST.searchHelper(rootNode.left, num)
	} else if rootNode.Data < num {
		return bST.searchHelper(rootNode.right, num)
	} else {
		return true
	}
}

func (bST BinarySearchTree) String() string {
	printBST(bST.node)
	return ""
}

func printBST(node *NodeBinarySearchTree) {
	if node == nil {
		return
	}
	if node.left != nil || node.right != nil {
		fmt.Printf("this one is parent node -> current value %v\n", *node)
	} else if node.left == nil && node.right == nil {
		fmt.Printf("this one is leaf -> current value %v\n", *node)
	}

	printBST(node.left)
	printBST(node.right)
}

func (bST *BinarySearchTree) Remove(num int) {
	if bST.Search(num) {
		bST.removeHelper(bST.node, num)
	} else {
		fmt.Println("node couldn't be found")
	}
}

func (bST *BinarySearchTree) removeHelper(
	rootNode *NodeBinarySearchTree,
	num int,
) *NodeBinarySearchTree {
	if rootNode == nil {
		return rootNode
	} else if num < rootNode.Data {
		rootNode.left = bST.removeHelper(rootNode.left, num)
	} else if num > rootNode.Data {
		rootNode.right = bST.removeHelper(rootNode.right, num)
	} else {
		if rootNode.left == nil && rootNode.right == nil {
			rootNode = nil
		} else if rootNode.right != nil {
			rootNode.Data = bST.successor(rootNode)
			rootNode.right = bST.removeHelper(rootNode.right, rootNode.Data)
		} else {
			rootNode.Data = bST.predecessor(rootNode)
			rootNode.left = bST.removeHelper(rootNode.left, rootNode.Data)
		}
	}
	return rootNode
}

func (bST *BinarySearchTree) predecessor(root *NodeBinarySearchTree) int {
	root = root.left

	for root.right != nil {
		root = root.right
	}
	return root.Data
}

func (bST *BinarySearchTree) successor(root *NodeBinarySearchTree) int {
	root = root.right
	for root.left != nil {
		root = root.left
	}
	return root.Data
}
