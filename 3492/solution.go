package solution

func maxContainers(n int, w int, maxWeight int) int {
	for i := n * n; i > 0; i-- {
		if i*w <= maxWeight {
			return i
		}
	}

	return 0
}
