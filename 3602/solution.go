package solution

func concatHex36(n int) string {
	toBase := func(x int, b int) string {
		res := ""

		for x != 0 {
			remain := x % b

			if 0 <= remain && remain <= 9 {
				res = string(rune(remain+48)) + res
			} else if 10 <= remain && remain <= 35 {
				res = string(rune(remain+55)) + res
			}

			x /= b
		}

		return res
	}

	return toBase(n*n, 16) + toBase(n*n*n, 36)
}
