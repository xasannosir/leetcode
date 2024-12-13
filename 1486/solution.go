package solution

func xorOperation(n int, start int) int {
	var sum = start

	for i, j := start+2, 1; j < n; i, j = i+2, j+1 {
		sum ^= i
	}

	return sum
}
