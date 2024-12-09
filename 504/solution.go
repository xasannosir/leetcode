package solution

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	var (
		isNegative bool
		result     string
	)

	if num < 0 {
		num = -num
		isNegative = true
	}

	for ; num > 0; num /= 7 {
		result += string(rune((num % 7) + 48))
	}
	
	runes := []rune(result)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	if isNegative {
		return "-" + string(runes)
	}
	return string(runes)
}
