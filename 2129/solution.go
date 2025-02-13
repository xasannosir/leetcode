package solution

func capitalizeTitle(title string) string {
	capitalize := func(word string) string {
		var result string

		if word[0] >= 97 && word[0] <= 122 {
			result += string(word[0] - 32)
		} else {
			result += string(word[0])
		}

		for i := 1; i < len(word); i++ {
			if word[i] >= 65 && word[i] <= 90 {
				result += string(word[i] + 32)
			} else {
				result += string(word[i])
			}
		}

		return result
	}

	lower := func(word string) string {
		var result string
		for i := 0; i < len(word); i++ {
			if word[i] >= 65 && word[i] <= 90 {
				result += string(word[i] + 32)
			} else {
				result += string(word[i])
			}
		}

		return result
	}

	var words []string
	var result string

	var word string
	for i := 0; i < len(title); i++ {
		if title[i] != ' ' {
			word += string(title[i])
		} else {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		}
	}

	if len(word) > 0 {
		words = append(words, word)
	}

	for _, v := range words {
		if len(v) > 2 {
			result += capitalize(v) + " "
		} else {
			result += lower(v) + " "
		}
	}

	return result[:len(result)-1]
}
