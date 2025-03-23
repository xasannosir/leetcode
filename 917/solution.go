package solution

func reverseOnlyLetters(s string) string {
	var letters []byte

	for i := 0; i < len(s); i++ {
		if (65 <= s[i] && s[i] <= 90) || (97 <= s[i] && s[i] <= 122) {
			letters = append(letters, s[i])
		}
	}

	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	index := 0
	var result string

	for i := 0; i < len(s); i++ {
		if (65 <= s[i] && s[i] <= 90) || (97 <= s[i] && s[i] <= 122) {
			result += string(letters[index])
			index++
		} else {
			result += string(s[i])
		}
	}

	return result
}
