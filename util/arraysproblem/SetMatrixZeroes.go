package arraysproblem

func setZeroes(matrix [][]int) {
	hMap := make(map[int][]int)
	counter := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				hMap[counter] = []int{i, j}
				counter += 1
			}
		}
	}
	for _, v := range hMap {
		horizontal := v[0]
		vertical := v[1]

		matrix[horizontal] = make([]int, len(matrix[horizontal]))
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if j == vertical {
					matrix[i][j] = 0
				}
			}
		}
	}
}

func SetZeroes(matrix [][]int) {
	setZeroes(matrix)
}
