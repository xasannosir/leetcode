package solution

func greatestLetter(s string) string {
	upper, lower := "", ""

	for i := 0; i < len(s); i++ {
		if 65 <= s[i] && s[i] <= 90 {
			upper += string(s[i])
		} else {
			lower += string(s[i])
		}
	}

	sort := func(word string) string {
		bytes := []byte(word)

		for i := 0; i < len(bytes); i++ {
			for j := 0; j < len(bytes)-i-1; j++ {
				if bytes[j] > bytes[j+1] {
					bytes[j], bytes[j+1] = bytes[j+1], bytes[j]
				}
			}
		}

		return string(bytes)
	}

	upper = sort(upper)
	lower = sort(lower)

	for i := len(upper) - 1; i >= 0; i-- {
		for j := len(lower) - 1; j >= 0; j-- {
			if upper[i]+32 == lower[j] {
				return string(upper[i])
			}
		}
	}

	return ""
}
