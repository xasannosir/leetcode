package solution

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	grid := make([][]int, 0)

	for i := 0; i < n; i++ {
		arr := make([]int, 0)
		for j := 0; j < m; j++ {
			arr = append(arr, matrix[j][i])
		}
		grid = append(grid, arr)
	}

	return grid
}
