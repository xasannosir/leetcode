package solution

func findAndReplacePattern(words []string, pattern string) []string {
	replaceNum := func(val string) int {
		counter := make(map[byte]int)

		for j := 0; j < len(val); j++ {
			counter[val[j]]++
		}

		number := 0

		for j := 0; j < len(val); j++ {
			number = number*10 + (counter[val[j]])
		}

		return number
	}

	checkOrder := func(pat, word string) bool {
		for i := 0; i < len(word); i++ {
			p, w := pat[i], word[i]
			for j := i; j < len(word); j++ {
				if pat[j] == p {
					if word[j] == w {
						continue
					} else {
						return false
					}
				}
			}
		}

		return true
	}

	numbers := make([]int, 0)

	for i := 0; i < len(words); i++ {
		numbers = append(numbers, replaceNum(words[i]))
	}

	search := replaceNum(pattern)

	var result []string
	for ind, num := range numbers {
		if num == search && checkOrder(pattern, words[ind]) {
			result = append(result, words[ind])
		}
	}

	return result
}
