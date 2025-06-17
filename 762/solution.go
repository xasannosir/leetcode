package solution

func countPrimeSetBits(left int, right int) int {
	total := 0

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}

		if n == 2 {
			return true
		}

		if n%2 == 0 {
			return false
		}

		for i := 3; i < n/2+1; i += 2 {
			if n%i == 0 {
				return false
			}
		}

		return true
	}

	for i := left; i <= right; i++ {
		count := 0
		num := i

		for num != 0 {
			count += num % 2
			num /= 2
		}

		if isPrime(count) {
			total++
		}
	}

	return total
}
