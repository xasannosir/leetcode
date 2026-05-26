package solution

func numberOfSpecialChars(word string) int {
	upper := make(map[rune]struct{})
	lower := make(map[rune]struct{})
	count := 0

	for _, w := range word {
		if w >= 65 && w <= 90 {
			lower[w] = struct{}{}
		} else {
			upper[w] = struct{}{}
		}
	}

	for l := range lower {
		if _, ok := upper[l+32]; ok {
			count++
		}
	}

	return count
}
