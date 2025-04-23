package solution

func countLargestGroup(n int) int {
	var sumOfNum func(num int) int

	sumOfNum = func(num int) int {
		if num > 9 {
			return num%10 + sumOfNum(num/10)
		}

		return num
	}

	counter := make(map[int]int)

	for i := 1; i <= n; i++ {
		counter[sumOfNum(i)]++
	}

	maxLen := 1

	for _, v := range counter {
		maxLen = max(maxLen, v)
	}

	count := 0
	for _, v := range counter {
		if v == maxLen {
			count++
		}
	}

	return count
}
