package solution

func equalPairs(grid [][]int) int {
	isEqualArray := func(arr1 []int, arr2 []int) bool {
		if len(arr1) != len(arr2) {
			return false
		}

		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				return false
			}
		}

		return true
	}
	var rows [][]int
	var cols [][]int
	count := 0

	for i := 0; i < len(grid); i++ {
		rows = append(rows, grid[i])
	}

	for i := 0; i < len(grid); i++ {
		var col []int
		for j := 0; j < len(grid[i]); j++ {
			col = append(col, grid[j][i])
		}
		cols = append(cols, col)
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if isEqualArray(rows[i], cols[j]) {
				count++
			}
		}
	}

	return count
}
