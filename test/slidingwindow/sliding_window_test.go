package slidingwindow

import (
	"fmt"
	"project/util/slidingwindow"
	"testing"
)

func TestMaximumAverageSubarray1(t *testing.T) {
	input := []int{1,12,-5,-6,50,3}
	// input := []int{-10,-8,-3,1}
	k := 4
	fmt.Println(slidingwindow.MaximumAverageSubarray1(input, k))
}