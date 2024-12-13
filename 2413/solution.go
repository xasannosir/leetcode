package solution

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}

	var start = n + 1
	for {
		if start%2 == 0 && start%n == 0 {
			return start
		}
		start += 2
	}
}
