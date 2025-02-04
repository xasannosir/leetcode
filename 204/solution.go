package solution

func countPrimes(n int) int {
	primes := make([]bool, n)

	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for p := 2; p*p < n; p++ {
		if primes[p] == true {
			for i := p * p; i < n; i += p {
				primes[i] = false
			}
		}
	}

	count := 0
	for p := 2; p < n; p++ {
		if primes[p] {
			count++
		}
	}

	return count
}
