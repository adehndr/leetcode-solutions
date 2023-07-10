package test

import (
	"fmt"
	"project/algo"
	"project/algo/search"
	"project/algo/sorting"
	"project/data"
	"testing"
)

func TestDFSAlgo(t *testing.T) {
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

	algo.DepthFirstSearch(&gM, 2)
}

func TestBFSAlgo(t *testing.T) {
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
	algo.BreadthFirstSearch(&gM, 1)
	algo.BreadthFirstSearchOptimized(&gM, 1)
}

func TestBFSAlgoOptimized(t *testing.T) {
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
	algo.BreadthFirstSearchOptimized(&gM, 1)
}

func TestMergeSort(t *testing.T) {
	input := []int{4,3,2,1}
	sorting.MergeSort(input)
	fmt.Println(input)
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1,2,3,4,5,6}
	fmt.Println(search.BinarySearch(nums,-3))
}
