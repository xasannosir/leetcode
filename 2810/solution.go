package solution

func finalString(s string) string {
	reverse := func(val string) string {
		bytes := []byte(val)

		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}

		return string(bytes)
	}

	var result string

	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			result = reverse(result)
		} else {
			result += string(s[i])
		}
	}

	return result
}
