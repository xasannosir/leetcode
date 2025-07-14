package solution

func deleteGreatestValue(grid [][]int) int {
	sort := func(nums []int) {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums)-i-1; j++ {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		sort(grid[i])
	}

	total := 0

	for len(grid[0]) != 0 {
		maxVal := 0
		for i := 0; i < len(grid); i++ {
			maxVal = max(maxVal, grid[i][len(grid[i])-1])
			grid[i] = grid[i][:len(grid[i])-1]
		}
		total += maxVal
	}

	return total
}
