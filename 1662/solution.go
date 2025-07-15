package solution

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	join := func(words []string) string {
		result := ""

		for _, word := range words {
			result += word
		}

		return result
	}

	return join(word1) == join(word2)
}
