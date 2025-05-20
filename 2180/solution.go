package solution

func countEven(num int) int {
	var sum func(n int) int

	sum = func(n int) int {
		if n < 10 {
			return n
		}

		return n%10 + sum(n/10)
	}

	count := 0
	for i := 2; i <= num; i++ {
		if sum(i)%2 == 0 {
			count++
		}
	}

	return count
}
