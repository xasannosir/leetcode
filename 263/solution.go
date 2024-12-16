package solution

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	isPrime := func(n int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	if isPrime(n) {
		if !(n == 2 || n == 3 || n == 5) {
			return false
		}
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			if isPrime(i) {
				if !(i == 2 || i == 3 || i == 5) {
					return false
				}
			}
			if isPrime(n / i) {
				if !(n/i == 2 || n/i == 3 || n/i == 5) {
					return false
				}
			}
		}
	}

	return true
}
