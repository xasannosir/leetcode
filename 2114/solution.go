package solution

func mostWordsFound(sentences []string) int {
	max := 0

	split := func(s string) []string {
		var (
			result []string
			word   string
		)

		for i := 0; i < len(s); i++ {
			if s[i] != ' ' {
				word += string(s[i])
			} else {
				result = append(result, word)
				word = ""
			}
		}
		result = append(result, word)

		return result
	}

	for _, value := range sentences {
		lst := split(value)

		if len(lst) > max {
			max = len(lst)
		}
	}

	return max
}
