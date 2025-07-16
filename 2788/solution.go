package solution

func splitWordsBySeparator(words []string, separator byte) []string {
	res := make([]string, 0)

	split := func(s string, sep byte) []string {
		var result []string
		var word string

		for i := 0; i < len(s); i++ {
			if s[i] != sep {
				word += string(s[i])
			} else {
				if word != "" {
					result = append(result, word)
					word = ""
				}
			}
		}
		if word != "" {
			result = append(result, word)
		}

		return result
	}

	for _, word := range words {
		res = append(res, split(word, separator)...)
	}

	return res
}
