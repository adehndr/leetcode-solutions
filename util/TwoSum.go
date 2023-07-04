package util

func TwoSums(nums []int, target int) []int{
	hMap := map[int]int{}
	for k, v := range nums {
		complement := target - v
		if idx, ok := hMap[complement]; ok {
			return []int{idx,k}
		}else {
			hMap[v] = k
		}
	}
	return []int{0,0}
}