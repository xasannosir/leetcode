package solution

func maximumValue(strs []string) int {
	onlyDigit := func(s string) bool {
		for i := 0; i < len(s); i++ {
			if !(48 <= s[i] && s[i] <= 57) {
				return false
			}
		}

		return true
	}

	toNum := func(s string) int {
		res := 0

		for i := 0; i < len(s); i++ {
			res = res*10 + int(s[i]-48)
		}

		return res
	}

	maxVal := 0

	for _, str := range strs {
		if onlyDigit(str) {
			maxVal = max(maxVal, toNum(str))
		} else {
			maxVal = max(maxVal, len(str))
		}
	}

	return maxVal
}
