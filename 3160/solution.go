package solution

func queryResults(limit int, queries [][]int) []int {
	balls := make(map[int]int)
	colors := make(map[int]int)
	result := make([]int, len(queries))

	for index, query := range queries {
		x, y := query[0], query[1]

		colors[y]++
		if c, ok := balls[x]; ok {
			colors[c]--
			if colors[c] == 0 {
				delete(colors, c)
			}
		}
		balls[x] = y
		result[index] = len(colors)
	}

	return result
}
