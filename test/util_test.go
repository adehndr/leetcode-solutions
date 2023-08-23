package test

import (
	"fmt"
	"project/util"
	"strings"
	"testing"
)

func TestFibonnaciRecursive(t *testing.T) {
	util.FibonnaciRecursive(5)
}

func TestCompareCharacter(t *testing.T) {
	str := "ADE"
	a := str[1]
	b := strings.ToLower(string(a))
	fmt.Println("D" == string(str[1]))
	fmt.Println(b)

}
