package solution

func sortSentence(s string) string {
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

	sorted := func(s []string) []string {
		for i := 0; i < len(s); i++ {
			for j := i; j < len(s); j++ {
				if s[i][len(s[i])-1] > s[j][len(s[j])-1] {
					temp := s[i]
					s[i] = s[j]
					s[j] = temp
				}
			}
		}

		return s
	}

	join := func(s []string) string {
		var result string

		for i := 0; i < len(s); i++ {
			result += string(s[i][:len(s[i])-1]) + " "
		}

		return result[:len(result)-1]
	}

	return join(sorted(split(s)))
}
