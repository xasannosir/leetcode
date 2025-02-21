package solution

import "sort"

func sortVowels(s string) string {
	isVowel := func(b byte) bool {
		switch b {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		default:
			return false
		}
	}

	sortedVowels := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			sortedVowels = append(sortedVowels, s[i])
		}
	}

	sort.Slice(sortedVowels, func(i, j int) bool {
		return sortedVowels[i] < sortedVowels[j]
	})

	increase := 0
	res := make([]byte, 0, len(sortedVowels))

	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			res = append(res, sortedVowels[increase])
			increase++
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}
