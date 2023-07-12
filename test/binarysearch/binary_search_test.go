package binarysearch

import (
	"fmt"
	"testing"
	"project/util/binarysearch"
)

func TestFindSmallestGreaterThanTarget(t *testing.T) {
	tmp := []string{"e","e","e","e","e","e","n","n","n","n"}
	target := "e"
	tmp2 := []byte{}
	for _, v := range tmp {
		tmp2 = append(tmp2, v[0])
	}
	fmt.Println(tmp2, target, target[0])

	fmt.Println(string(binarysearch.FindSmallestGreaterTarget(tmp2, target[0])))
}
