package solution

func sumOfNumberAndReverse(num int) bool {
	reverse := func(n int) int {
		res := 0

		for n != 0 {
			res = res*10 + n%10
			n /= 10
		}

		return res
	}

	for i := 0; i <= num; i++ {
		if i+reverse(i) == num {
			return true
		}
	}

	return false
}
