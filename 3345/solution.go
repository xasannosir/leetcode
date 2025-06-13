package solution

func smallestNumber(n int, t int) int {
	sumOfDigits := func(num int) int {
		total := 1

		for num != 0 {
			total *= num % 10
			num /= 10
		}

		return total
	}

	for {
		if sumOfDigits(n)%t == 0 {
			return n
		}

		n++
	}
}
