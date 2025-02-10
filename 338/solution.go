package solution

func countBits(n int) []int {
	var result []int

	for i := 0; i <= n; i++ {
		temp := i
		count := 0

		for temp >= 2 {
			count += temp % 2
			temp /= 2
		}

		result = append(result, count+temp)
	}

	return result
}
