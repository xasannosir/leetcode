package solution

func lexicalOrder(n int) []int {
	res := make([]int, 0)
	var next func(num int)

	next = func(num int) {
		num *= 10

		for i := 0; i < 10; i++ {
			if num <= n {
				res = append(res, num)
				next(num)
			} else {
				break
			}

			num++
		}
	}

	for i := 1; i < 10; i++ {
		if i <= n {
			res = append(res, i)
			next(i)
		}
	}

	return res
}
