package twopointers

func PatternString(subStr, fullString string) int {
	leftPtr := 0
	rightPtr := 0
	res := 0
	for rightPtr < len(fullString) {
		if string(fullString[rightPtr]) == string(subStr[leftPtr]) {
			leftPtr += 1
		} else {
			leftPtr = 0
		}

		if leftPtr == len(subStr) {
			res += 1
			leftPtr = 0
		}
		rightPtr += 1
	}
	return res
}