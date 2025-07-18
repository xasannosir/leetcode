package solution

func replaceWords(dictionary []string, sentence string) string {
	hasPrefix := func(s string, sep string) bool {
		if len(sep) > len(s) {
			return false
		}

		for i := 0; i < len(s) && i < len(sep); i++ {
			if s[i] != sep[i] {
				return false
			}
		}

		return true
	}

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

	words := split(sentence)
	res := ""

	for _, word := range words {
		isHave := false
		finds := make([]string, 0)
		for _, dic := range dictionary {
			if hasPrefix(word, dic) {
				finds = append(finds, dic)
				isHave = true
			}
		}

		if !isHave {
			res += word + " "
		} else {
			shortest := finds[0]

			for i := 1; i < len(finds); i++ {
				if len(finds[i]) < len(shortest) {
					shortest = finds[i]
				}
			}

			res += shortest + " "
		}
	}

	return res[:len(res)-1]
}
