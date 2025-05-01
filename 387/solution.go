package solution

func firstUniqChar(s string) int {
	findIndex := func(search byte, str string) int {
		for i := 0; i < len(str); i++ {
			if str[i] == search {
				return i
			}
		}

		return -1
	}

	appear := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		appear[s[i]]++
	}

	index := -1
	for val, count := range appear {
		if count == 1 && index == -1 {
			index = findIndex(val, s)
		} else if count == 1 {
			index = min(index, findIndex(val, s))
		}
	}

	return index
}
