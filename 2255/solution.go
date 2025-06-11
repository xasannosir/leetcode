package solution

func countPrefixes(words []string, s string) int {
	count := 0

	isPrefix := func(s string, sub string) bool {
		if s == sub {
			return true
		}

		if len(sub) > len(s) {
			return false
		}

		for i := 0; i < len(s) && i < len(sub); i++ {
			if s[i] != sub[i] {
				return false
			}
		}

		return true
	}

	for _, word := range words {
		if isPrefix(s, word) {
			count++
		}
	}

	return count
}
