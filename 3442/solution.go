package solution

func maxDifference(s string) int {
	freq := make(map[byte]int)
	a1, a2 := 0, len(s)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	for _, v := range freq {
		if v%2 == 1 {
			a1 = max(a1, v)
		} else if v%2 == 0 {
			a2 = min(a2, v)
		}
	}

	return a1 - a2
}
