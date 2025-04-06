package solution

func countOdds(low int, high int) int {
	count := 0

	if low%2 == 0 {
		low++
	}

	for i := low; i <= high; i += 2 {
		count++
	}

	return count
}
