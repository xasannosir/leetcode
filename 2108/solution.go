package solution

func firstPalindrome(words []string) string {
	isPalindrome := func(s string) bool {
		bytes := []byte(s)

		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if bytes[i] != bytes[j] {
				return false
			}
		}

		return true
	}

	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}

	return ""
}
