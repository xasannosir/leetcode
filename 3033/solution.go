package solution

func modifiedMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == -1 {
				maxVal := 0
				for k := 0; k < len(matrix); k++ {
					maxVal = max(maxVal, matrix[k][j])
				}
				matrix[i][j] = maxVal
			}
		}
	}

	return matrix
}
