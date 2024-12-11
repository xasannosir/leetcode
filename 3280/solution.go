package solution

func convertDateToBinary(date string) string {
	var arr = []string{date[0:4], date[5:7], date[8:10]}
	var result string

	for _, value := range arr {
		var integer int
		for i := 0; i < len(value); i++ {
			integer = integer*10 + int(value[i]-48)
		}

		var str string
		for integer > 0 {
			str += string(rune((integer % 2) + 48))
			integer /= 2
		}

		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		result += string(runes) + "-"
	}

	return result[:len(result)-1]
}
