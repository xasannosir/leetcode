package solution

func getNoZeroIntegers(n int) []int {
	start := 1

	hasZero := func(num int) bool {
		for num != 0 {
			if num%10 == 0 {
				return true
			}
			num /= 10
		}

		return false
	}

	for {
		if !hasZero(start) && !hasZero(n-start) {
			return []int{start, n - start}
		}
		start++
	}
}
