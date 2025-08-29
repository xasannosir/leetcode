package solution

func lengthOfLongestSubstring(s string) int {
	counter := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	maxVal := len(counter)

	for maxVal != 0 {
		for i := 0; i+maxVal <= len(s); i++ {
			checker := make(map[byte]int)

			for j := i; j < i+maxVal; j++ {
				checker[s[j]]++
			}

			if len(checker) == maxVal {
				return maxVal
			}
		}

		maxVal--
	}

	return 0
}
