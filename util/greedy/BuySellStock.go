package greedy

func BuySellStock(prices []int) int {
	var tmpResult int
	if len(prices) <= 1 {
		return tmpResult
	}

	l := 0
	r := 1
	// [7,1,5,3,6,4]
	for r < len(prices) {
		if prices[l] > prices[r] {
			l = r
		} else {
			diff := prices[r] - prices[l]
			if diff > tmpResult {
				tmpResult = diff
			}
		}
		r++
	}
	return tmpResult
}
