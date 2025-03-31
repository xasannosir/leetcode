package solution

func selfDividingNumbers(left int, right int) []int {
	isDiv := func(num int) bool {
		temp := num

		for temp > 0 {
			if temp%10 == 0 {
				return false
			} else if num%(temp%10) != 0 {
				return false
			}
			temp /= 10
		}

		return true
	}

	var res []int

	for i := left; i <= right; i++ {
		if isDiv(i) {
			res = append(res, i)
		}
	}

	return res
}
