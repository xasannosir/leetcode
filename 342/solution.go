package solution

func isPowerOfFour(n int) bool {
	var start = 1

	for ; start <= n; start *= 4 {
		if start == n {
			return true
		}
	}

	return false
}
