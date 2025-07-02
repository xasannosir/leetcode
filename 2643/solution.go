package solution

func rowAndMaximumOnes(mat [][]int) []int {
	row, count := 0, 0

	for i := 0; i < len(mat); i++ {
		curCount := 0
		for j := 0; j < len(mat[i]); j++ {
			curCount += mat[i][j]
		}
		if curCount > count {
			count = curCount
			row = i
		}
	}

	return []int{row, count}
}
