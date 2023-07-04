package util

func FibonnaciRecursive(n int) int {
	if n <= 2 {
		return 1
	}
	return FibonnaciRecursive(n-2) + FibonnaciRecursive(n-1)
}

func FibonnaciMemoize(n int, memo []int) int {
	var result int
	if n <= 2 {
		result = 1
	} else {
		result = FibonnaciMemoize(n-2, memo) + FibonnaciMemoize(n-1, memo)
	}
	memo[n] = result
	return memo[n]
}

func FibonnacoMemoizeBottomUp(n int) int {
	memo := make([]int, n+1)

	if n <= 1 {
		return 1
	}

	memo[0] = 1
	memo[1] = 1

	for k := 2; k < len(memo); k++ {
		memo[k] = memo[k-2] + memo[k-1]
	}
	return memo[n]
}
