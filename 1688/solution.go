package solution

func numberOfMatches(n int) int {
	count := 0

	for {
		if n%2 == 0 {
			count += n / 2
			n /= 2
		} else {
			count += (n - 1) / 2
			n = (n-1)/2 + 1
		}

		if n == 1 {
			break
		}
	}

	return count
}
