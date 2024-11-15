package solution

func generate(numRows int) [][]int {
	var result = make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1

		for j := 1; j < i+1; j++ {
			if !(j == 0 || j == i) {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}

	return result
}
