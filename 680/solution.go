package solution

func validPalindrome(s string) bool {
	reverse := func(str string) string {
		r := []rune(str)

		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}

		return string(r)
	}

	if s == reverse(s) {
		return true
	}

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			if string(runes[:i])+string(runes[i+1:]) == reverse(string(runes[:i])+string(runes[i+1:])) {
				return true
			} else if string(runes[:j])+string(runes[j+1:]) == reverse(string(runes[:j])+string(runes[j+1:])) {
				return true
			} else {
				return false
			}
		}
	}

	return false
}
