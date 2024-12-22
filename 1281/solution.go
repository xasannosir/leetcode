package solution

func subtractProductAndSum(n int) int {
	product, sum := 1, 0

	for n > 0 {
		latest := n % 10
		product *= latest
		sum += latest
		n /= 10
	}

	return product - sum
}
