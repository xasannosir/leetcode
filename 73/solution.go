package solution

func setZeroes(matrix [][]int) {
	var temp [][]int

	for i := 0; i < len(matrix); i++ {
		var arr []int
		for j := 0; j < len(matrix[i]); j++ {
			arr = append(arr, matrix[i][j])
		}
		temp = append(temp, arr)
	}

	fill := func(k int, l int) {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if i == k || j == l {
					matrix[i][j] = 0
				}
			}
		}
	}

	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(temp[i]); j++ {
			if temp[i][j] == 0 {
				fill(i, j)
			}
		}
	}
}
