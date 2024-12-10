package solution

func isPowerOfThree(n int) bool {
	var start = 1

	for ; start <= n; start *= 3 {
		if start == n {
			return true
		}
	}

	return false
}
