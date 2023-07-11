package dynamicprogramming

func climbingStairs(n int) int {
	tempRes := []int{1, 1}
	fibWithMemoization(tempRes, 3)
	return 0
}

func fibSimpleRec(target int) int {
	if target <= 1 {
		return 1
	}
	return fibSimpleRec(target-1) + fibSimpleRec(target-2)
}

func fibWithMemoization(memo []int, n int) int {
	var result int

	if n <= 2 {
		result = 1
	} else {
		result = fibWithMemoization(memo, n-1) + fibWithMemoization(memo, n-2)
	}
	return result

}

func fibWithMemoizationBottomTop(n int) int {
	temp := []int{}
	if n <= 1 {
		return 1
	}
	temp = append(temp, 1)
	temp = append(temp, 1)

	for i := 2; i <= n; i++ {
		temp = append(temp, temp[i-1]+temp[i-2])
	}
	return temp[len(temp)-1]
}
