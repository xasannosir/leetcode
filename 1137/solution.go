package solution

func tribonacci(n int) int {
	var t = []int{0, 1, 1}

	if n < 3 {
		return t[n]
	}

	for i := 3; i < n+1; i++ {
		var sum = 0
		for j := 0; j < len(t); j++ {
			sum += t[j]
		}

		t[0] = t[1]
		t[1] = t[2]
		t[2] = sum
	}

	return t[2]
}
