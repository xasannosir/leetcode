package solution

func prefixCount(words []string, pref string) int {
	count := 0

	for _, word := range words {
		if len(word) >= len(pref) {
			isPref := true
			for i := 0; i < len(pref); i++ {
				if word[i] != pref[i] {
					isPref = false
					break
				}
			}

			if isPref {
				count++
			}
		}
	}

	return count
}
