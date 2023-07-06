package arraysproblem

import (
	"fmt"
	"project/util/arraysproblem"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	fmt.Println(arraysproblem.ProductExceptSelf([]int{1, 2, 3, 4}))
}

func TestProductOfArrayExceptSelf2(t *testing.T) {
	fmt.Println(arraysproblem.ProductExceptSelf2([]int{1, 2, 3, 4}))
}

func TestAccessingSlice(t *testing.T) {
	aNum := []int{1, 2, 3, 4}

	fmt.Println(recAccessSlice(aNum[2:]))
}

func recAccessSlice(slice []int) int {
	if len(slice) == 0 {
		return 1
	}
	return slice[0] * recAccessSlice(slice[1:])
}
