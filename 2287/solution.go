package solution

func rearrangeCharacters(s string, target string) int {
	count := 0
	freq := make(map[rune]int)
	isEnd := false
	targetMap := make(map[rune]int)

	for _, val := range target {
		targetMap[val]++
	}

	for _, val := range s {
		freq[val]++
	}

	for {
		for key, val := range targetMap {
			if count, ok := freq[key]; ok && count >= val {
				freq[key] -= val
			} else {
				delete(freq, key)
				isEnd = true
				break
			}
		}

		if !isEnd {
			count++
			isEnd = false
		}

		try := false
		for key, val := range targetMap {
			if count, ok := freq[key]; ok && count >= val {
				try = true
				break
			}
		}

		if !try {
			break
		}
	}

	return count
}
