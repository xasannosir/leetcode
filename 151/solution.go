package solution

func reverseWords(s string) string {
	var words []string
	var word string
	var result string

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		}
	}

	if len(word) > 0 {
		words = append(words, word)
	}

	for i := len(words) - 1; i >= 0; i-- {
		result += words[i] + " "
	}

	return result[:len(result)-1]
}
