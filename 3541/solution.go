package solution

func maxFreqSum(s string) int {
	isVowel := func(character byte) bool {
		switch character {
		case 'a':
			return true
		case 'e':
			return true
		case 'i':
			return true
		case 'o':
			return true
		case 'u':
			return true
		default:
			return false
		}
	}

	counter := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	vowelMaxFreq := 0
	consonantMaxFreq := 0

	for symbol, count := range counter {
		if isVowel(symbol) {
			vowelMaxFreq = max(vowelMaxFreq, count)
		} else {
			consonantMaxFreq = max(consonantMaxFreq, count)
		}
	}

	return consonantMaxFreq + vowelMaxFreq
}
