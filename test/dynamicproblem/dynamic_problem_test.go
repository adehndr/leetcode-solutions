package test

import (
	"fmt"
	"project/util/dynamicprogramming"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{2}
	amount := 3

	fmt.Println(dynamicprogramming.CoinChange(coins, amount))
}

func TestMaximumSubarray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(dynamicprogramming.MaximumSubarray(input))
}
