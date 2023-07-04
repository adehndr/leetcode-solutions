package util

import (
	"strings"
)

func BackspaceString(s, t string) bool {
	aStc := Stack{}
	bStc := Stack{}
	for _, v := range s {
		if string(v) == "#" {
			aStc.Pop()
		} else {
			aStc.Push(string(v))
		}
	}

	for _, v := range t {
		if string(v) == "#" {
			bStc.Pop()
		} else {
			bStc.Push(string(v))
		}
	}

	if strings.Join([]string(aStc), "") == strings.Join([]string(bStc), "") {
		return true
	}
	return false
}
