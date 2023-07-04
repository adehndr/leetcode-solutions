package data

import "fmt"

type LinkedList struct {
	Head *Node
}

func (lL *LinkedList) AddNode(inputNode *Node) {

	if lL.Head == nil {
		lL.Head = inputNode
		return
	}

	currHead := lL.Head

	for currHead.Tail != nil {
		currHead = currHead.Tail
	}

	currHead.Tail = inputNode
}

func (lL *LinkedList) Remove(anInput string) {
	if lL.Head == nil {
		return
	}

	if lL.Head.Data == anInput {
		lL.Head = lL.Head.Tail
		return
	}

	currHead := lL.Head

	for currHead.Tail != nil && currHead.Tail.Data != anInput {
		currHead = currHead.Tail
	}

	if currHead.Tail != nil {
		currHead.Tail = currHead.Tail.Tail
	}

}

func (lL *LinkedList) Traverse() {
	currHead := lL.Head

	for currHead != nil {
		fmt.Printf("%s -> ", currHead.Data)
		currHead = currHead.Tail
	}
	fmt.Println()
}

func (lL *LinkedList) GetNode(idx int) *Node {
	init := 0

	currHead := lL.Head

	for currHead != nil {
		if init == idx {
			return currHead
		}
		init += 1
		currHead = currHead.Tail
	}
	return nil
}
