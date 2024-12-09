package main

func isHappy(n int) bool {
	if n == 1 || n == 7 {
		return true
	}
	for n > 9 {
		var sum = 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if sum == 1 || sum == 7 {
			return true
		}
		n = sum
	}

	return false
}
