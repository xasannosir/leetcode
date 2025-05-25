package solution

func countDigits(num int) int {
	var digits func(n int, count int) (int, int)

	digits = func(n int, count int) (int, int) {
		if num%(n%10) == 0 {
			count++
		}

		if n < 10 {
			return n, count
		}

		return digits(n/10, count)
	}

	_, count := digits(num, 0)

	return count
}
