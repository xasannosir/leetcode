package solution

func reverseVowels(s string) string {
	var vowelIndices []int

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowelIndices = append(vowelIndices, i)
		} else if s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			vowelIndices = append(vowelIndices, i)
		}
	}

	runes := []rune(s)
	for i, j := 0, len(vowelIndices)-1; i < j; i, j = i+1, j-1 {
		runes[vowelIndices[i]], runes[vowelIndices[j]] = runes[vowelIndices[j]], runes[vowelIndices[i]]
	}

	return string(runes)
}
