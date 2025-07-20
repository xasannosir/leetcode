package solution

func checkDivisibility(n int) bool {
	product := func(num int) int {
		result := 1

		for num != 0 {
			result *= num % 10
			num /= 10
		}

		return result
	}

	sum := func(num int) int {
		result := 0

		for num != 0 {
			result += num % 10
			num /= 10
		}

		return result
	}

	return n%(sum(n)+product(n)) == 0
}
