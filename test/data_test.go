package test

import (
	"fmt"
	"project/data"
	"testing"
)

func TestLinkedList(t *testing.T) {
	lList := data.LinkedList{Head: &data.Node{Data: "A"}}
	lList.AddNode(&data.Node{Data: "B"})
	lList.AddNode(&data.Node{Data: "C"})
	lList.AddNode(&data.Node{Data: "D"})

	lList.Traverse()
	lList.Remove("C")
	fmt.Println(lList.GetNode(6).Data)
	lList.Traverse()
}

func TestQueue(t *testing.T) {
	aQ := data.QueueString{}
	aQ.EnqueueString("A")
	aQ.EnqueueString("B")
	aQ.EnqueueString("C")
	aQ.EnqueueString("D")

	fmt.Println(aQ)
	aQ.DequeueString()
	aQ.DequeueString()
	aQ.DequeueString()
	aQ.DequeueString()
	aQ.DequeueString()

	fmt.Println(aQ)
}

func TestGraphMatrix(t *testing.T) {
	gM := data.GraphMatrix{}
	gM.InitMatrix(5)
	gM.AddNode(&data.Node{Data: "A"})
	gM.AddNode(&data.Node{Data: "B"})
	gM.AddNode(&data.Node{Data: "C"})
	gM.AddNode(&data.Node{Data: "D"})
	gM.AddNode(&data.Node{Data: "E"})

	gM.AddEdge(0, 1)
	gM.AddEdge(1, 2)
	gM.AddEdge(1, 4)
	gM.AddEdge(2, 3)
	gM.AddEdge(2, 4)
	gM.AddEdge(4, 0)
	gM.AddEdge(4, 2)

	gM.Print()
}
