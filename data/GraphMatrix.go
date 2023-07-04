package data

import "fmt"

type GraphMatrix struct {
	Matrix [][]int
	Nodes  LinkedList
}

func (gM *GraphMatrix) InitMatrix(size int) {
	for i := 0; i < size; i++ {
		(*gM).Matrix = append((*gM).Matrix, []int{})
		for j := 0; j < size; j++ {
			(*gM).Matrix[i] = append((*gM).Matrix[i], 0)
		}
	}
}

func (gM *GraphMatrix) Print() {
	fmt.Print("  ")
	for i := range (*gM).Matrix {
		fmt.Printf("%s ", (*gM).Nodes.GetNode(i).Data)
	}
	fmt.Println()

	for iIdx, iValue := range (*gM).Matrix {
		fmt.Printf("%s ", (*gM).Nodes.GetNode(iIdx).Data)
		for _, jValue := range iValue {
			fmt.Printf("%d ", jValue)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (gM *GraphMatrix) AddEdge(src, dst int) {
	(*gM).Matrix[src][dst] = 1
}

func (gM *GraphMatrix) AddNode(node *Node) {
	gM.Nodes.AddNode(node)
}

func (gM *GraphMatrix) CheckEdge(src, dst int) bool {
	if (*gM).Matrix[src][dst] == 1 {
		return true
	}
	return false
}
