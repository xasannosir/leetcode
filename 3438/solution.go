package solution

func findValidPair(s string) string {
	counter := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] && counter[s[i]] == int(s[i]-48) && counter[s[i+1]] == int(s[i+1]-48) {
			return string(s[i]) + string(s[i+1])
		}
	}

	return ""
}
