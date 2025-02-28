package solution

func totalMoney(n int) int {
	total, start, stop, num := 0, 1, 7, 1

	for i := 0; i < n; i++ {
		total += num
		num++
		if num > stop {
			start++
			stop++
			num = start
		}
	}

	return total
}
