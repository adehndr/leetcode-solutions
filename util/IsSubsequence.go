package util

import (
	"fmt"
	"strings"
)

func IsSubsequence(s string, t string) bool {
	findSlice := strings.Split(s, "")
	textSlice := strings.Split(t, "")
	hMap := map[string][]int{}

	for k, v := range textSlice {
		for _, f := range findSlice {
			if v == f {
				if _, ok := hMap[v]; ok {
					temp := hMap[v]
					temp = append(temp, k)
					hMap[v] = temp
				} else {
					hMap[v] = []int{k}
				}
			}
		}
	}
	foundedSlice := []int{}
	tmpNonDupliMap := map[int]string{}
	for k1, v := range hMap {
		for _, v2 := range v {
			tmpNonDupliMap[v2] = k1
		}
	}
	fmt.Println(tmpNonDupliMap, hMap)
	for k := range tmpNonDupliMap {
		foundedSlice = append(foundedSlice, k)
	}
	mergeSortA(foundedSlice)

	var constuctedStr string
	for _, v := range foundedSlice {
		constuctedStr += textSlice[v]
	}
	fmt.Println(constuctedStr)
	if constuctedStr == s || strings.Contains(constuctedStr, s) {
		return true
	} else {
		return false
	}
}

func mergeSortA(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := len(nums) / 2
	tmpLeft := []int{}
	tmpRight := []int{}
	tmpLeft = append(tmpLeft, nums[0:mid]...)
	tmpRight = append(tmpRight, nums[mid:]...)

	mergeSortA(tmpLeft)
	mergeSortA(tmpRight)
	merge(nums, tmpLeft, tmpRight)
}

func merge(arr, left, right []int) {
	l := 0
	r := 0
	i := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			arr[i] = left[l]
			i += 1
			l += 1
		} else {
			arr[i] = right[r]
			i += 1
			r += 1
		}
	}

	for l < len(left) {
		arr[i] = left[l]
		i += 1
		l += 1
	}

	for r < len(right) {
		arr[i] = right[r]
		i += 1
		r += 1
	}
}
