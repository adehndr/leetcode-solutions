package util

func MajorityElement(nums []int) int {
	hMap := make(map[int]int)

	majority := len(nums) / 2
	for _, v := range nums {
		if _, ok := hMap[v]; ok {
			tmp := hMap[v]
			tmp += 1
			hMap[v] = tmp
		} else {
			hMap[v] = 1
		}

		if value, ok := hMap[v]; ok {
			if value > majority {
				return v
			}
		}
	}

	return -1
}
