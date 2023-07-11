package intervals

func MergeIntervals(intervals [][]int) [][]int {
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	result := [][]int{}

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][len(intervals[i])-1] >= intervals[i+1][0] {
			tempRes := []int{intervals[i][0], intervals[i+1][len(intervals[i+1])-1]}
			result = append(result, tempRes)
			i += 1
		} else {
			result = append(result, intervals[i])
		}
	}

	if len(intervals) > 1 && len(result) < len(intervals)-1 && intervals[len(intervals)-2][len(intervals[len(intervals)-2])-1] >= intervals[len(intervals)-1][0] {
		result = append(result, []int{intervals[len(intervals)-2][0], intervals[len(intervals)-1][len(intervals[len(intervals)-1])-1]})
	} else if len(result) < len(intervals)-1 {
		result = append(result, intervals[len(intervals)-1])
	}

	return result
}

func merge2(interval [][]int) [][]int {
	for {
		
	}
}