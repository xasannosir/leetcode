package solution

func myAtoi(s string) int {
	res := 0
	isPositive := true
	digitStart := false
	isCut := false
	cutIndex := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			isCut = true
			cutIndex = i
		} else {
			break
		}
	}

	if isCut {
		s = s[cutIndex+1:]
	}

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		isPositive = false
		s = s[1:]
	} else if s[0] == '+' {
		isPositive = true
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		if 48 <= s[i] && s[i] <= 57 {
			if s[i] != '0' {
				res = res*10 + int(s[i]-48)
				digitStart = true
			} else if digitStart {
				res = res*10 + int(s[i]-48)
			}
		} else {
			break
		}

		if isPositive && res > 2_147_483_647 {
			res = 2_147_483_647
		} else if !isPositive && res > 2_147_483_648 {
			res = 2_147_483_648
		}
	}

	if !isPositive {
		res = -res
	}

	return res
}
