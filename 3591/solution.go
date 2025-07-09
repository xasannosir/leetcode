package solution

func checkPrimeFrequency(nums []int) bool {
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}

		if num == 2 {
			return true
		}

		if num%2 == 0 {
			return false
		}

		for i := 3; i <= num/2; i += 2 {
			if num%i == 0 {
				return false
			}
		}

		return true
	}

	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	for _, v := range freq {
		if isPrime(v) {
			return true
		}
	}

	return false
}
