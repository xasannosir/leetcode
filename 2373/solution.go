package solution

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	maxLocal := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		arr := make([]int, 0)
		for j := 0; j < n-2; j++ {
			maxVal := 0
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					maxVal = max(maxVal, grid[k][l])
				}
			}
			arr = append(arr, maxVal)
		}
		maxLocal = append(maxLocal, arr)
	}

	return maxLocal
}
