package solution

func getLucky(s string, k int) int {
	num := 0

	for i := 0; i < len(s); i++ {
		digit := int(s[i] - 96)

		num += (digit / 10) + (digit % 10)
	}

	for i := 1; i < k; i++ {
		total := 0

		for num != 0 {
			total += num % 10
			num /= 10
		}

		num = total

		if num <= 9 {
			break
		}
	}

	return num
}
