package solution

func hammingWeight(n int) int {
	count := 0

	for n >= 2 {
		count += n % 2
		n /= 2
	}

	return count + n
}
