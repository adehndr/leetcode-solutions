package twopointers

func MoveZeroesBruteForce(nums []int) {
	// 0,1,2,3
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] == 0 {
				tmp := nums[j+1]
				nums[j+1] = 0
				nums[j] = tmp
			}
		}
	}
}

func MoveZeroes(nums []int) {
	moveZeroesNeetCode(nums)
}

func moveZeroesNeetCode(nums []int) {
	l := 0
	// Utilizing l as pointer to know the last zero
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			tmp := nums[r]
			nums[r] = nums[l]
			nums[l] = tmp
			l += 1
		}
	}

}
