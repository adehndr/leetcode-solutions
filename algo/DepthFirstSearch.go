package algo

import (
	"fmt"
	"project/data"
)

func DepthFirstSearch(gM *data.GraphMatrix, src int) {
	visited := make([]bool, 5)
	helperDFS(gM, src, visited)
}

func helperDFS(gM *data.GraphMatrix, src int, visited []bool) {
	if visited[src] {
		return
	} else {
		visited[src] = true
		fmt.Printf("%s = visited\n", (*gM).Nodes.GetNode(src).Data)
	}
	for i := 0; i < len((*gM).Matrix[src]); i++ {
		if (*gM).Matrix[src][i] == 1 {
			helperDFS(gM, i, visited)
		}
	}
}
