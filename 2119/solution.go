package solution

func isSameAfterReversals(num int) bool {
	reverse := func(n int) int {
		res := 0

		for n != 0 {
			res = res*10 + n%10
			n /= 10
		}

		return res
	}

	return num == reverse(reverse(num))
}
