package solution

func maximumWealth(accounts [][]int) int {
	var maximum int

	for _, account := range accounts {
		var sum int
		for _, val := range account {
			sum += val
		}
		if sum > maximum {
			maximum = sum
		}
	}

	return maximum
}
