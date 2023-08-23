package backtracingproblem

import (
	"fmt"
	"project/util/backtracking"
	"testing"
)

func TestWordSearch(t *testing.T) {
	input := [][]byte{}
	input = append(input, []byte("ABCE"))
	input = append(input, []byte("SFCS"))
	input = append(input, []byte("ADEE"))
	fmt.Println(backtracking.WordSearch(input, "ABCCED"))
}

func TestIterateString(t *testing.T) {
	a := "sebuah"
	for _, chr := range a {
		fmt.Println(string(chr))
	}
	// create arr
	input := [][]int{{1,2,3},{4,5,6}}
	another := make([][]bool,len(input))
	for i := 0; i < len(input); i++ {
		another[i] = make([]bool, len(input[i]))
	}
	fmt.Println(input)
	fmt.Println(another)
}
