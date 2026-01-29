package solution

func countCharacters(words []string, chars string) int {
	freq := make(map[byte]int)
	length := 0

	for i := 0; i < len(chars); i++ {
		freq[chars[i]]++
	}

	for _, word := range words {
		item := make(map[byte]int)
		end := true

		for j := 0; j < len(word); j++ {
			item[word[j]]++
		}

		for k, v := range item {
			count, ok := freq[k]

			if (ok && count < v) || !ok {
				end = false
				break
			}
		}

		if end {
			length += len(word)
		}
	}

	return length
}
