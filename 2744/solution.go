package solution

func maximumNumberOfStringPairs(words []string) int {
	reverse := func(s string) string {
		bytes := []byte(s)

		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}

		return string(bytes)
	}

	count := 0

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i] == reverse(words[j]) {
				count++
			}
		}
	}

	return count
}
