package solution

func isCircularSentence(sentence string) bool {
	split := func(s string) []string {
		var (
			words []string
			word  string
		)

		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				words = append(words, word)
				word = ""
			} else {
				word += string(s[i])
			}
		}

		words = append(words, word)

		return words
	}

	words := split(sentence)

	for i := 0; i < len(words)-1; i++ {
		if words[i][len(words[i])-1] != words[i+1][0] {
			return false
		}
	}

	return words[0][0] == words[len(words)-1][len(words[len(words)-1])-1]
}
