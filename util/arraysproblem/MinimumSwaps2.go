package arraysproblem

// This problem was coming from hackerrank

func MinimumSwaps2(arr []int32) int32 {
	return minimumSwaps2(arr)
}

func minimumSwaps2(arr []int32) int32 {
	var minSwap int32 = 0
	hMap := make(map[int32]int, len(arr))

	for k, v := range arr {
		hMap[v] = k
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != int32(i+1) {
			correctPosition := hMap[int32(i+1)]
			tempVal := arr[i]
			arr[i] = arr[correctPosition]
			arr[correctPosition] = tempVal
			minSwap += 1
			hMap[int32(i+1)] = i
			hMap[int32(tempVal)] = correctPosition
		}
	}

	return minSwap
}
