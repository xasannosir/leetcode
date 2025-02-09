package solution

func convertToTitle(columnNumber int) string {
	var result string

	for columnNumber > 0 {
		columnNumber = columnNumber - 1
		result = string(rune(columnNumber%26+65)) + result
		columnNumber = columnNumber / 26
	}

	return result
}
