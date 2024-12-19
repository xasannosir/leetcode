package solution

func diagonalSum(mat [][]int) int {
	var sum int

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if i == j || i+j == len(mat[i])-1 {
				sum += mat[i][j]
			}
		}
	}

	return sum
}
