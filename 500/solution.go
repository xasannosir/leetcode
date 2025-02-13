package solution

func findWords(words []string) []string {
	var result []string

	firstRow := "qwertyuiopQWERTYUIOP"
	secondRow := "asdfghjklASDFGHJKL"
	thirdRow := "zxcvbnmZXCVBNM"

	inRow := func(symbol byte, row string) bool {
		for i := 0; i < len(row); i++ {
			if row[i] == symbol {
				return true
			}
		}

		return false
	}

	checkAll := func(word string, row string) bool {
		isComplete := true

		for i := 0; i < len(word); i++ {
			if !inRow(word[i], row) {
				isComplete = false
				break
			}
		}

		return isComplete
	}

	for _, word := range words {
		if inRow(word[0], firstRow) && checkAll(word, firstRow) {
			result = append(result, word)
		} else if inRow(word[0], secondRow) && checkAll(word, secondRow) {
			result = append(result, word)
		} else if inRow(word[0], thirdRow) && checkAll(word, thirdRow) {
			result = append(result, word)
		} else {
			continue
		}
	}

	return result
}
