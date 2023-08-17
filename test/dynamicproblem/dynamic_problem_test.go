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
