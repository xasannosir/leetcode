package solution

func replaceDigits(s string) string {
	var result []byte

	result = append(result, s[0])
	for i := 1; i < len(s); i++ {
		if 48 <= s[i] && s[i] <= 57 {
			sign := byte(int(s[i]) + int(s[i-1]-48))
			result = append(result, sign)
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}
