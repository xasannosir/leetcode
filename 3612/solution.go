package solution

func processStr(s string) string {
	reverse := func(val string) string {
		bytes := []byte(val)

		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}

		return string(bytes)
	}

	result := ""

	for i := 0; i < len(s); i++ {
		if 97 <= s[i] && s[i] <= 122 {
			result += string(s[i])
		} else if s[i] == '*' && len(result) > 0 {
			result = result[:len(result)-1]
		} else if s[i] == '#' {
			result += result
		} else if s[i] == '%' {
			result = reverse(result)
		}
	}

	return result
}
