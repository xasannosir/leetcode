package solution

func freqAlphabets(s string) string {
	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '#' {
			result = string(s[i]+48) + result
		} else {
			result = string((s[i-2]-48)*10+(s[i-1]-48)+96) + result
			i -= 2
		}
	}

	return result
}
