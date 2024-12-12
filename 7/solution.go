package solution

func reverse(x int) int {
	var result int

	for x != 0 {
		result = result*10 + (x % 10)
		x = x / 10
	}

	var (
		maxInt32 = 2147483647
		minInt32 = -2147483648
	)

	if result > maxInt32 || result < minInt32 {
		return 0
	}

	return result
}
