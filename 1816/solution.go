package solution

func truncateSentence(s string, k int) string {
	var words []string

	var word string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else {
			words = append(words, word)
			word = ""
		}
	}
	words = append(words, word)

	var result string

	for i := 0; i < k; i++ {
		if i < len(words) {
			result += words[i] + " "
		} else {
			break
		}
	}

	return result[:len(result)-1]
}
