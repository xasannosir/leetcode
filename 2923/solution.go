package solution

func findChampion(grid [][]int) int {
	indexCount := make(map[int]int)

	for i := 0; i < len(grid); i++ {
		sum := 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
		}
		indexCount[i] = sum
	}

	strongTeam := 0
	strongBalance := 0

	for key, value := range indexCount {
		if strongBalance < value {
			strongBalance = value
			strongTeam = key
		}
	}

	return strongTeam
}
