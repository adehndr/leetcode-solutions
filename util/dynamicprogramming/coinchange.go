package dynamicprogramming

import "fmt"

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}

func coinChange(coins []int, amount int) int {
	dpArr := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dpArr[i] = amount + 1
	}

	dpArr[0] = 0

	for i := 1; i < amount+1; i++ {
		for _, eachC := range coins {
			if i-eachC >= 0 {
				dpArr[i] = min(dpArr[i], 1+dpArr[i-eachC])
			}
		}
	}
	fmt.Println(dpArr)
	if dpArr[len(dpArr)-1] == amount + 1 {
		return -1
	} else {
		return dpArr[len(dpArr)-1]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
