package solution

func maximumOddBinaryNumber(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			res += string(s[i])
		} else {
			res = string(s[i]) + res
		}
	}

	return res[1:] + string(res[0])
}
