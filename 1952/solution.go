package solution

func isThree(n int) bool {
	count := 0

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count == 1
}
