package util

import "fmt"

func FindSmallestGreaterTarget(letters []byte, target byte) byte {
	diffWithIndex := make(map[int]byte)
	idxSmallest := 0
	var smallest int
	for k, v := range letters {
		if target < v {
			diffWithIndex[k] = v - (target)
			smallest = int(v - (target))
			idxSmallest = k
		}
	}
	fmt.Println("target", target, "smallest", smallest)
	for k, v := range diffWithIndex {
		if v <= byte(smallest) {
			smallest = int(v)
			idxSmallest = k
		}
	}

	fmt.Println(diffWithIndex, smallest, idxSmallest , letters[idxSmallest])

	return letters[idxSmallest]
}
