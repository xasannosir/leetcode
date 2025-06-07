package solution

func smallestRepunitDivByK(k int) int {
	n := 0

	for i := 0; i < k; i++ {
		n = (n*10 + 1) % k

		if n == 0 {
			return i + 1
		}
	}

	return -1
}
