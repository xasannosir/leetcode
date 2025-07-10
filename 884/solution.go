package solution

func uncommonFromSentences(s1 string, s2 string) []string {
	split := func(s string) []string {
		var result []string
		var word string

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

	counter1 := make(map[string]int)
	counter2 := make(map[string]int)
	res := make([]string, 0)

	words1 := split(s1)
	words2 := split(s2)

	for _, word := range words1 {
		counter1[word]++
	}

	for _, word := range words2 {
		counter2[word]++
	}

	for k, v := range counter1 {
		if _, ok := counter2[k]; !ok && v == 1 {
			res = append(res, k)
		} else {
			delete(counter1, k)
			delete(counter2, k)
		}
	}

	for k, v := range counter2 {
		if _, ok := counter1[k]; !ok && v == 1 {
			res = append(res, k)
		} else {
			delete(counter2, k)
			delete(counter1, k)
		}
	}

	return res
}
