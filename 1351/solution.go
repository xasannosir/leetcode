package solution

func countNegatives(grid [][]int) int {
	var count int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}

	return count
}
