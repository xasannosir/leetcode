package solution

func maxNumberOfBalloons(text string) int {
	count := 0
	freq := make(map[rune]int)
	isEnd := false

	ballon := map[rune]int{
		'b': 1,
		'a': 1,
		'l': 2,
		'o': 2,
		'n': 1,
	}

	for _, val := range text {
		freq[val]++
	}

	for {
		for key, val := range ballon {
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
		for key, val := range ballon {
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
