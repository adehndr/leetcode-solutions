package algo

import (
	"fmt"
	"project/data"
)

func BreadthFirstSearch(gM *data.GraphMatrix, src int) {
	visited := make([]bool, len((*gM).Matrix))

	aQueue := data.QueueInteger{}
	aQueue.EnqueueInteger(src)

	for !aQueue.IsEmptyInteger() {
		fmt.Println(aQueue)
		newSrc := aQueue.DequeueInteger()
		visited[newSrc] = true
		fmt.Printf("Current src %d dengan visited %v\n", newSrc, visited)
		fmt.Printf("%s = visited\n", (*gM).Nodes.GetNode(newSrc).Data)

		for i := 0; i < len((*gM).Matrix[newSrc]); i++ {
			fmt.Printf("i ke %d apakah sama dengan 1 ? %v || udah dikunjungin %v ? || queue mengandung element ? %v\n", i, (*gM).Matrix[newSrc][i] == 1, visited[i], aQueue.Contains(i))
			if (*gM).Matrix[newSrc][i] == 1 && !visited[i] && !aQueue.Contains(i) {
				aQueue.EnqueueInteger(i)
			}
		}
	}

}

func BreadthFirstSearchOptimized(gM *data.GraphMatrix, src int) {
	aQueue := data.QueueInteger{}
	visited := make([]bool, len((*gM).Matrix))

	aQueue.EnqueueInteger(src)
	visited[src] = true

	for !aQueue.IsEmptyInteger() {
		src = aQueue.DequeueInteger()
		fmt.Printf("%s = visited\n", (*gM).Nodes.GetNode(src).Data)

		for i := 0; i < len((*gM).Matrix[src]); i++ {
			if (*gM).Matrix[src][i] == 1 && !visited[i] {
				aQueue.EnqueueInteger(i)
				visited[i] = true
			}
		}

	}
}
