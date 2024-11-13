package solution

func climbStairs(n int) int {
	var a, b = 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
