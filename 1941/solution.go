package solution

func areOccurrencesEqual(s string) bool {
	counter := make(map[byte]int)
	checker := make(map[int]struct{})

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	for _, v := range counter {
		checker[v] = struct{}{}

		if len(checker) > 1 {
			return false
		}
	}

	return true
}
