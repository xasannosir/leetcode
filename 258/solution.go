package solution

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	var result int

	for {
		for num > 0 {
			result += num % 10
			num /= 10
		}

		if result < 10 {
			break
		} else {
			num = result
			result = 0
		}
	}

	return result
}
