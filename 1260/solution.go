package solution

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	arr := make([]int, n*m)
	count := k % (n * m)

	for i, c := 0, 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[c] = grid[i][j]
			c++
		}
	}

	for i := 0; i < count; i++ {
		for j := n*m - 1; j >= 1; j-- {
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
		}
	}

	for i, c := 0, 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = arr[c]
			c++
		}
	}

	return grid
}
