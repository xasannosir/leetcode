package solution

func findLucky(arr []int) int {
	counter := make(map[int]int)
	lucky := -1

	for _, num := range arr {
		counter[num]++
	}

	for k, v := range counter {
		if k == v && k > lucky {
			lucky = k
		}
	}

	return lucky
}
