package solution

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, val := range row {
			if val == target {
				return true
			}
		}
	}

	return false
}
