package solution

func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	} else if divisor == -1 {
		if dividend < 0 && -dividend > 2147483647 {
			return 2147483647
		} else {
			return -dividend
		}
	}

	var (
		isNegative = false
		res        int
	)

	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		isNegative = true
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	for dividend >= divisor {
		dividend -= divisor
		res++
	}

	if isNegative {
		return -res
	} else {
		if res > 2147483647 {
			return 2147483647
		} else {
			return res
		}
	}
}
