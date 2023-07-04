package util

func MoveZeroes(nums []int) {
	aMap := map[int][]int{}
	// 0 , 1 , 2
	for idx, e := range nums {
		if e == 0 {
			aMap[e] = append(aMap[e], idx)
		}
	}
}

func MoveZeroesBruteForce(nums []int){
	// 0,1,2,3
	for i := 0; i < len(nums) - 1; i++ {
		for j := 0; j < len(nums) - 1; j++ {
			if nums[j] == 0 {
				tmp := nums[j+1]
				nums[j+1] = 0
				nums[j] = tmp
			}
		}
	}

}
