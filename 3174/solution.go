package solution

func clearDigits(s string) string {
	var result string

	for i := 0; i < len(s); i++ {
		if 97 <= s[i] && s[i] <= 122 {
			result += string(s[i])
		} else if 48 <= s[i] && s[i] <= 57 && len(result) != 0 {
			result = result[:len(result)-1]
		}
	}

	return result
}
