package solution

func vowelStrings(words []string, left int, right int) int {
	count := 0

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

	for i := left; i <= right; i++ {
		if isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			count++
		}
	}

	return count
}
