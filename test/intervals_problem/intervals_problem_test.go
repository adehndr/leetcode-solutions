package intervalsproblem

import (
	"fmt"
	"project/util/intervals"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	input := [][]int{}
	input = append(input, []int{1,3})
	input = append(input, []int{2,6})
	input = append(input, []int{8,10})
	input = append(input, []int{15,18})

/* 	input = append(input, []int{1,4})
	input = append(input, []int{4,5}) */

	fmt.Println(intervals.MergeIntervals(input))
}