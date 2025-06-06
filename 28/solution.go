package solution

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && len(haystack[i:]) >= len(needle) {
			isHave := true
			ind := i
			for j := 0; j < len(needle); j++ {
				if haystack[ind] != needle[j] {
					isHave = false
					break
				}
				ind++
			}

			if isHave {
				return i
			}
		}
	}

	return -1
}
