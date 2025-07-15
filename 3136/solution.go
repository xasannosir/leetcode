package solution

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	vowel, consonant := false, false

	isVowel := func(s byte) bool {
		switch s {
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
		case 'A':
			return true
		case 'E':
			return true
		case 'I':
			return true
		case 'O':
			return true
		case 'U':
			return true
		default:
			return false
		}
	}

	for i := 0; i < len(word); i++ {
		if 65 <= word[i] && word[i] <= 90 || 97 <= word[i] && word[i] <= 122 {
			if isVowel(word[i]) {
				vowel = true
			} else {
				consonant = true
			}
		} else if !(48 <= word[i] && word[i] <= 57) {
			return false
		}
	}

	return vowel && consonant
}
