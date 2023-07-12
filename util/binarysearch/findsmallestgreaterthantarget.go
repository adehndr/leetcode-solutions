package binarysearch

import "fmt"

func findSmallestGreaterTarget(letters []byte, target byte) byte {
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
	fmt.Println("target", target, "smallest", smallest, "diffwithindex", diffWithIndex)
	for k, v := range diffWithIndex {
		if v <= byte(smallest) {
			smallest = int(v)
			idxSmallest = k
		}
	}

	fmt.Println(diffWithIndex, smallest, idxSmallest, letters[idxSmallest])

	return letters[idxSmallest]
}

func FindSmallestGreaterTarget(letters []byte, target byte) byte {
	return findSmallestGreaterTarget2(letters, target)
}

func findSmallestGreaterTarget2(letters []byte, target byte) byte {
	ptr := -1

	for k, v := range letters {
		if v > target {
			ptr = k
			break
		}
	}
	if ptr == -1 {
		return letters[0]
	} else {
		return letters[ptr]
	}
}

func findSmallestGreaterTarget3(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	if letters[left] > target || letters[right] <= target {
		return letters[left]
	}

	mid := left + ((right - left) / 2)
	fmt.Println("left right mid", left, right, mid)
	for left+1 < right {
		mid = left + ((right - left) / 2)
		fmt.Println("left right mid", left, right, mid)
		if target >= letters[mid] {
			left = mid
		} else {
			right = mid
		}
	}

	return letters[right]

}
