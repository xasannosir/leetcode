package solution

func isPowerOfTwo(n int) bool {
	var start = 1

	for ; start <= n; start *= 2 {
		if start == n {
			return true
		}
	}

	return false
}
